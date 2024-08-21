package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// 如何验证优雅的重启？
// 1. 打开终端，执行 go build -o gin_graceful_restart gin_graceful_restart.go 编译程序，然后执行 ./gin_graceful_restart 终端会输出当前的 pid，例如：[pid] 12345
// 2. 将代码中处理请求中的 message 原来为 pong 现在改成为 pong1
// 3. 再次执行 go build -o gin_graceful_restart gin_graceful_restart.go 编译程序
// 4. 打开浏览器，访问 http://127.0.0.1:8080/ping 此时浏览器应该会白屏等待服务端返回响应
// 5. 打开另一个终端，执行 kill -1 12345 命令给程序发送 syscall.SIGHUP 信号，其中 12345 为第 1 步中的 pid
// 6. 依旧在第 4 步的浏览器中等待，等响应到 pong 信息之后，再次刷新页面，此时应该会看到响应为 pong1 信息
// 7. 这样就在不影响当前未处理完请求的同时完成了程序代码的替换，从而实现了优雅重启

// 但是需要注意的是：此时程序的 PID 变化了，因为 endless 是通过 fork 子进程处理新请求，待原进程处理完当前请求后再退出的方式实现优雅重启的。
// 所以当你的项目是使用类似 supervisor 的软件管理进程时就不适用这种方式了。

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// 模拟程序处理请求需要 5 秒
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认 endless 服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM 和 syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发 `fork/restart` 实现优雅重启（kill -1 pid 会发送 SIGHUP 信号）
	// 接收到 syscall.SIGINT 或 syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发 HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的 hook 函数
	if err := endless.ListenAndServe(":8080", router); err != nil {
		log.Printf("Server err: %v\n", err)
	}

	log.Println("Server exiting")
}
