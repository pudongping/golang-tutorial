package main

import (
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

var (
	ErrTokenGenFailed         = errors.New("令牌生成失败")
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrTokenNotFound          = errors.New("无法找到令牌")
)

// JWT 定义一个 jwt 对象
type JWT struct {

	// 密钥，用以加密 JWT
	Key []byte

	// 刷新 token 的最大过期时间（单位：分钟）（必须大于 ExpireTime 这里的用途就是说只要当前时间没有超过 MaxRefresh 设置的时间，都允许通过旧 token 来换取新 token）
	MaxRefresh int64

	// 定义 token 过期时间（单位：分钟）即当颁发 token 后，多少分钟后 token 过期
	ExpireTime int64

	// token 的签发者
	Issuer string
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	// user_id 是自定义的字段
	UserID string `json:"user_id"` // 当前登录的用户 id

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtPkg.StandardClaims
}

func NewJWT(secret, issuer string, maxRefreshTime, expireTime int64) *JWT {
	if maxRefreshTime <= expireTime {
		log.Fatal("最大刷新时间必须大于 token 的过期时间")
	}

	return &JWT{
		Key:        []byte(secret), // 密钥
		MaxRefresh: maxRefreshTime, // 允许刷新时间
		ExpireTime: expireTime,     // token 过期时间
		Issuer:     issuer,         // token 的签发者
	}
}

// ParseToken 解析 token
func (j *JWT) ParseToken(c *gin.Context, userToken ...string) (*JWTCustomClaims, error) {
	var (
		tokenStr string
		err      error
	)

	if len(userToken) > 0 {
		tokenStr = userToken[0]
	} else {
		// 获取 token
		tokenStr, err = j.GetToken(c)
		if err != nil {
			return nil, err
		}
	}

	// 解析用户 token
	token, err := j.parseTokenString(tokenStr)

	// 解析出错时
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			switch validationErr.Errors {
			case jwtPkg.ValidationErrorMalformed:
				return nil, ErrTokenMalformed
			case jwtPkg.ValidationErrorExpired:
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	// Valid 验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），
	// 需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// GetTTL 计算出 token 还剩多少秒后过期
func (j *JWT) GetTTL(c *gin.Context, userToken ...string) (int64, error) {

	claims, err := j.ParseToken(c, userToken...)

	if err != nil {
		// 此时已经过期，或者出现 token 解析失败
		return 0, err
	}

	// 此时的 token 一定是没有过期的，否则上一步 ParseToken 就已经报错了
	ttl := claims.ExpiresAt - time.Now().Unix()

	return ttl, nil
}

// RefreshToken 刷新 token
func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 获取 token
	tokenStr, err := j.GetToken(c)
	if err != nil {
		return "", err
	}

	// 解析用户 token
	token, err := j.parseTokenString(tokenStr)

	// 解析出错时（未报错证明是合法的 token 或者未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 如果满足刷新 token 的条件，就继续往下走下一步（只要是单一的 ValidationErrorExpired 报错就认为是）
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 解析出自定义的载荷信息 JWTCustomClaims
	claims := token.Claims.(*JWTCustomClaims)

	maxRefreshTime := time.Duration(j.MaxRefresh) * time.Minute

	// 检查是否过了【最大允许刷新的时间】
	// 首次签名时间 + 最大允许刷新时间区间 > 当前时间 ====> 首次签名时间 > 当前时间 - 最大允许刷新时间区间
	if claims.IssuedAt > time.Now().Add(-maxRefreshTime).Unix() {
		// 此时并没有过最大允许刷新时间，因此可以重新颁发 token
		// 在这里重新赋值一下过期时间 ExpiresAt 从而达到刷新 token 的目的
		// 但是需要注意的是：一定不能更改 IssuedAt，因为这个字段是用来判断是否过了最大允许刷新时间的
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}

	// 当前时间过了最大允许刷新的时间
	// 因此就必须要重新登录了
	return "", ErrTokenExpiredMaxRefresh
}

// GenerateToken 生成 token
func (j *JWT) GenerateToken(userId string) (string, error) {
	// 构造用户 claims 信息（负荷）
	expireAtTime := j.expireAtTime()
	now := time.Now().Unix()
	claims := JWTCustomClaims{
		UserID: userId,
		StandardClaims: jwtPkg.StandardClaims{
			NotBefore: now,          // 签名生效时间
			IssuedAt:  now,          // 首次签名时间（后续刷新 token 不会更新）
			ExpiresAt: expireAtTime, // 签名过期时间
			Issuer:    j.Issuer,     // 签名颁发者
		},
	}

	// 根据 claims 生成 token
	token, err := j.createToken(claims)
	if err != nil {
		log.Printf("generate token failed: %v \n", err)
		return "", ErrTokenGenFailed
	}

	return token, nil
}

// createToken 创建 token，用于内部调用
func (j *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用 HS256 算法生成的 token
	tokenClaims := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	// 生成签名字符串
	return tokenClaims.SignedString(j.Key)
}

// expireAtTime 获取过期时间点
func (j *JWT) expireAtTime() int64 {
	expire := time.Duration(j.ExpireTime) * time.Minute
	// 返回加过期时间区间后的时间点
	return time.Now().Add(expire).Unix()
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (j *JWT) parseTokenString(tokenStr string) (*jwtPkg.Token, error) {
	// ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *jwtPkg.Token
	return jwtPkg.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return j.Key, nil
	})
}

// GetToken 获取请求中的 token 参数
func (j *JWT) GetToken(c *gin.Context) (string, error) {
	var token string

	if query, exists := c.GetQuery("token"); exists && "" != query {
		token = query
	} else if post, exists := c.GetPostForm("token"); exists && "" != post {
		token = post
	} else if authorization := c.GetHeader("Authorization"); "" != authorization {
		// 按空格进行分割
		parts := strings.SplitN(authorization, " ", 2)
		// 如果长度为 2 并且第一个元素为 Bearer
		if len(parts) == 2 && parts[0] == "Bearer" {
			// 返回第二个元素，即就是我们想要的 token
			token = parts[1]
		}
	} else {
		token = c.GetHeader("token")
	}

	if "" == token {
		return "", ErrTokenNotFound
	}

	return token, nil
}

func main() {
	c := gin.Default()

	c.POST("/signup", signUpHandler)
	c.GET("/me", AuthJWTMiddleware(), meHandler)

	if err := c.Run(":8080"); err != nil {
		log.Fatalf("run server failed: %v", err)
	}

}

const (
	// SecretKey 密钥
	SecretKey = "secret1"
	// Issuer 签发者
	Issuer = "project-name"
	// MaxRefreshTime 最大刷新时间
	MaxRefreshTime = 120
	// ExpireTime 过期时间
	ExpireTime = 10

	// CtxUserIDKey 上下文中用户 id 的 key
	CtxUserIDKey = "current_user_id"
)

type signUpParam struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func signUpHandler(c *gin.Context) {
	var param signUpParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if param.UserName != "admin" || param.Password != "123456" {
		c.JSON(400, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	// 生成 token
	j := NewJWT(SecretKey, Issuer, MaxRefreshTime, ExpireTime)
	token, err := j.GenerateToken("1688") // 假设当前用户的 id 为 1688
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		j := NewJWT(SecretKey, Issuer, MaxRefreshTime, ExpireTime)

		claims, err := j.ParseToken(c)
		if err != nil {
			msg := err.Error()
			if errors.Is(err, ErrTokenNotFound) {
				msg = "无法找到 token，请在请求头或者请求参数中携带 token 或者在请求头中携带 Authorization"
			}

			c.JSON(401, gin.H{
				"msg": msg,
			})
			c.Abort()
			return
		}

		spew.Dump(claims)

		// 将用户 id 存入上下文中
		c.Set(CtxUserIDKey, claims.UserID)
		c.Next()
	}
}

func meHandler(c *gin.Context) {
	userId := c.GetString(CtxUserIDKey)
	if userId == "" {
		c.JSON(401, gin.H{
			"msg": "未登录",
		})
		return
	}

	j := NewJWT(SecretKey, Issuer, MaxRefreshTime, ExpireTime)
	ttl, err := j.GetTTL(c)
	if err != nil {
		c.JSON(401, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"data": gin.H{
			"user_id":   userId,
			"token_ttl": ttl,
		},
	})
}
