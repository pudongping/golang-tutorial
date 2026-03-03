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
		"user_id":   "456",
		"email":     "delayed_user@example.com",
		"subject":   "Hello, Delayed World!",
		"send_time": time.Now().Format(time.DateTime),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("无法创建任务: %v", err)
	}

	task := asynq.NewTask(TypeEmailSend, data)
	// 入队任务，设置延迟30秒
	info, err := client.EnqueueContext(ctx, task,
		asynq.Queue(QueueDelayTaskEmailDefault), // 指定延迟任务专用队列
		asynq.MaxRetry(3),                       // 最大重试次数3次
		// 如果 ProcessIn 和 ProcessAt 同时出现时，则后面的会覆盖前面的延迟时间
		asynq.ProcessIn(30*time.Second), // 核心：延迟30秒后处理
		// asynq.ProcessAt(time.Now().Add(30*time.Second)), // 核心：延迟30秒后处理（也可以指定具体时间）
		asynq.Timeout(30*time.Second), // 任务处理超时时间30秒
	)
	if err != nil {
		log.Fatalf("无法入队任务: %v", err)
	}
	log.Printf("延迟任务已入队 ID=%s 队列=%s 将在 %v 后处理", info.ID, info.Queue, 30*time.Second)
}
