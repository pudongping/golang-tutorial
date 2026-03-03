package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// 创建一个 Asynq 客户端
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := map[string]string{
		"user_id":   "123",
		"email":     "user@example.com",
		"subject":   "Hello, World!",
		"send_time": time.Now().Format(time.DateTime),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("无法创建任务: %v", err)
	}

	task := asynq.NewTask(TypeEmailSend, data)
	info, err := client.EnqueueContext(
		ctx,
		task,
		asynq.Queue(QueueAsyncTaskEmailDefault), // 指定队列
		asynq.MaxRetry(3),                       // 最大重试次数3次
		asynq.Timeout(30*time.Second),           // 任务处理超时时间30秒
	)
	if err != nil {
		log.Fatalf("无法入队任务: %v", err)
	}
	log.Printf("任务已入队 ID=%s 队列=%s", info.ID, info.Queue)
}
