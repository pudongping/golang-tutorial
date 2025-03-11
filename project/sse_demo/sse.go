package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态页面路由（用于展示前端）
	r.StaticFile("/", "./index.html")

	// SSE 事件流路由
	r.GET("/events", func(c *gin.Context) {
		// 设置SSE必要响应头
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// 创建通道用于接收关闭通知
		clientClosed := c.Writer.CloseNotify()

		// 无限循环发送事件
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-clientClosed:
				fmt.Println("客户端断开连接")
				return
			case t := <-ticker.C:
				// SSE数据格式要求（重要！）
				event := fmt.Sprintf("data: %v\n\n", t.Format("2006-01-02 15:04:05"))

				// 发送数据到客户端
				c.SSEvent("message", event)
				c.Writer.Flush() // 立即刷新缓冲区
			}
		}
	})

	r.Run(":8080")
}
