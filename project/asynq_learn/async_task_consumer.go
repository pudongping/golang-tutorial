package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hibiken/asynq"
)

// HandleEmailSendTask 是处理发送邮件任务的函数
func HandleEmailSendTask(ctx context.Context, t *asynq.Task) error {
	spew.Dump("当前时间", time.Now().Format(time.DateTime))
	var payload map[string]string
	// 解析任务载荷
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal 失败: %v: %w", err, asynq.SkipRetry)
	}

	spew.Dump("正在处理发送邮件任务", payload)

	// 模拟发送邮件的耗时操作
	time.Sleep(2 * time.Second)

	log.Printf("邮件发送成功: email=%s", payload["email"])
	return nil
}

func main() {
	// 创建一个新的 Asynq 服务器实例
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB},
		asynq.Config{
			// 指定每个并发工人的数量
			Concurrency: 10,
			// 指定队列及其优先级
			Queues: map[string]int{
				QueueAsyncTaskEmailDefault: 10,
			},
			// 自定义错误处理程序
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Printf("任务处理失败: type=%s, payload=%s, err=%v", task.Type(), task.Payload(), err)
			}),
		},
	)

	// mux 用于将任务类型映射到处理程序
	mux := asynq.NewServeMux()

	// 注册任务处理函数
	mux.HandleFunc(TypeEmailSend, HandleEmailSendTask)

	// 启动服务器并阻塞等待信号
	log.Println("正在启动 Asynq 消费者服务器...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
