package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// 如何验证优雅的关机？
// 1. 打开终端，运行 go run gin_shutdown.go
// 2. 打开浏览器，并访问 http://127.0.0.1:8080/ping 此时浏览器应该会白屏等待服务端返回响应
// 3. 在刚刚打开的终端上迅速按下 Ctrl+C 命令，此时会自动给程序发送 syscall.SIGINT 信号
// 4. 此时程序并不会立即退出，而是会等上面的第 2 步的响应返回之后再退出，从而实现优雅关机的效果

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// 模拟程序处理请求需要 5 秒
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启一个 goroutine 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 服务启动失败
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个 5 秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 `Ctrl+C` 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个 5 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 5 秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
