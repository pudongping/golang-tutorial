package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit" // 令牌桶算法
	ratelimit1 "go.uber.org/ratelimit"     // 漏桶算法
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func pingHandler2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}

func rateLimit1() func(ctx *gin.Context) {
	// 漏桶算法，第一个参数为两滴水滴之间的时间间隔。
	// 此时表示两滴水之间的时间间隔是 100 纳秒
	rl := ratelimit1.New(100)

	return func(ctx *gin.Context) {
		// 尝试取出水滴
		if waitTime := rl.Take().Sub(time.Now()); waitTime > 0 {
			fmt.Printf("需要等待 %v 秒，下一滴水才会滴下来\n", waitTime)
			// 这里我们可以让程序继续等待，也可以直接拒绝掉
			// time.Sleep(waitTime)
			ctx.String(http.StatusOK, "rate limit, try again later")
			ctx.Abort()
			return
		}
		// 证明可以继续执行
		ctx.Next()
	}
}

func rateLimit2() func(ctx *gin.Context) {
	// 令牌桶算法：第一个参数为每秒填充令牌的速率为多少
	// 第二个参数为令牌桶的容量
	// 这里表示每秒填充 10 个令牌
	rl := ratelimit2.NewBucket(2*time.Second, 1)

	return func(ctx *gin.Context) {
		// 尝试取出令牌
		var num int64 = 1
		// 这里表示需要 num 个令牌和已经取出的令牌数是否相等
		// 不相等，则表示超过了限流
		// 比如，假设每一个请求过来消耗2个令牌，但是从桶中取出的令牌个数为 1 ，那么则认为超过了限流（一般而言是一个请求消耗一个令牌，这里仅为举例）
		if rl.TakeAvailable(num) != num {
			// 此次没有取到令牌，说明超过了限流
			ctx.String(http.StatusOK, "rate limit, try again later")
			ctx.Abort()
			return
		}
		// 证明可以继续执行
		ctx.Next()
	}
}

func main() {
	r := gin.Default()

	// 漏桶算法
	r.GET("/ping", rateLimit1(), pingHandler)

	// 令牌桶算法
	r.GET("/ping2", rateLimit2(), pingHandler2)

	r.Run()
}
