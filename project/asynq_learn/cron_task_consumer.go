package main

import (
	"context"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// HandleCronGreeting 处理定时问候任务
func HandleCronGreeting(ctx context.Context, t *asynq.Task) error {
	// 打印 "你好啊" + 当前时间（精确到秒）
	log.Printf("你好啊 %s \n", time.Now().Format(time.DateTime))
	return nil
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				QueueCronTask: 1, // 只监听定时任务队列
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Printf("任务处理失败: %v", err)
			}),
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeCronGreeting, HandleCronGreeting)

	log.Println("正在启动定时任务消费者...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
