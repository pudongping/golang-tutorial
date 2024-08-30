package main

import (
	"log"
	"time"

	jwtPkg "github.com/golang-jwt/jwt/v4"
)

type ARJWT struct {
	// 密钥，用以加密 JWT
	Key []byte

	// 定义 access token 过期时间（单位：分钟）即当颁发 access token 后，多少分钟后 access token 过期
	AccessExpireTime int64

	// 定义 refresh token 过期时间（单位：分钟）即当颁发 refresh token 后，多少分钟后 refresh token 过期
	// 一般来说，refresh token 的过期时间会比 access token 的过期时间长
	RefreshExpireTime int64

	// token 的签发者
	Issuer string
}

func NewARJWT(secret, issuer string, accessExpireTime, refreshExpireTime int64) *ARJWT {
	if refreshExpireTime <= accessExpireTime {
		log.Fatal("refresh token 过期时间必须大于 access token 过期时间")
	}
	return &ARJWT{
		Key:               []byte(secret),    // 密钥
		AccessExpireTime:  accessExpireTime,  // access token 过期时间
		RefreshExpireTime: refreshExpireTime, // refresh token 过期时间
		Issuer:            issuer,            // token 的签发者
	}
}

// GenerateToken 生成 access token 和 refresh token
func (arj *ARJWT) GenerateToken(userId string) (accessToken, refreshToken string, err error) {
	// 生成 access token 在 access token 中需要包含我们自定义的字段，比如用户 ID
	mc := JWTCustomClaims{
		UserID: userId,
		StandardClaims: jwtPkg.StandardClaims{
			// ExpiresAt 是一个时间戳，代表 access token 的过期时间
			ExpiresAt: time.Now().Add(time.Duration(arj.AccessExpireTime) * time.Minute).Unix(),
			// 签发人
			Issuer: arj.Issuer,
		},
	}

	// 生成 access token
	accessToken, err = jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, mc).SignedString(arj.Key)
	if err != nil {
		log.Printf("generate access token failed: %v \n", err)
		return "", "", err
	}

	// 生成 refresh token
	// refresh token 只需要包含标准的声明，不需要包含自定义的声明
	refreshToken, err = jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, jwtPkg.StandardClaims{
		// ExpiresAt 是一个时间戳，代表 refresh token 的过期时间
		ExpiresAt: time.Now().Add(time.Duration(arj.RefreshExpireTime) * time.Minute).Unix(),
		// 签发人
		Issuer: arj.Issuer,
	}).SignedString(arj.Key)

	return
}

func (arj *ARJWT) ParseAccessToken(tokenString string) (*JWTCustomClaims, error) {
	claims := new(JWTCustomClaims)

	token, err := jwtPkg.ParseWithClaims(tokenString, claims, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	})

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

	if _, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func (arj *ARJWT) RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 先判断 refresh token 是否有效
	if _, err = jwtPkg.Parse(refreshToken, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	}); err != nil {
		return
	}

	// 从旧的 access token 中解析出 JWTCustomClaims 数据出来
	var claims JWTCustomClaims
	_, err = jwtPkg.ParseWithClaims(accessToken, &claims, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	})
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 当 access token 是过期错误，并且 refresh token 没有过期时就创建一个新的 access token 和 refresh token
		if ok && validationErr.Errors == jwtPkg.ValidationErrorExpired {
			// 重新生成新的 access token 和 refresh token
			return arj.GenerateToken(claims.UserID)
		}
	}

	return
}
