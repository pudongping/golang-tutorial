package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 模拟批量发送任务，混合不同优先级的任务
	// 为了演示效果，我们多发一些任务
	taskCounts := 10

	// 1. 发送铂金会员任务 (最高优先级)
	for i := 0; i < taskCounts; i++ {
		enqueueTask(ctx, client, "user_platinum", QueuePriorityPlatinum, fmt.Sprintf("Platinum User %d", i))
	}

	// 2. 发送钻石会员任务 (中等优先级)
	for i := 0; i < taskCounts; i++ {
		enqueueTask(ctx, client, "user_diamond", QueuePriorityDiamond, fmt.Sprintf("Diamond User %d", i))
	}

	// 3. 发送普通会员任务 (低优先级)
	for i := 0; i < taskCounts; i++ {
		enqueueTask(ctx, client, "user_normal", QueuePriorityNormal, fmt.Sprintf("Normal User %d", i))
	}

	log.Println("所有任务入队完成，请启动消费者观察处理顺序")
}

func enqueueTask(ctx context.Context, client *asynq.Client, userID, queueName, subject string) {
	payload := map[string]string{
		"user_id": userID,
		"email":   fmt.Sprintf("%s@example.com", userID),
		"subject": subject,
		"queue":   queueName,
	}
	data, _ := json.Marshal(payload)
	task := asynq.NewTask(TypeEmailSend, data)

	info, err := client.EnqueueContext(ctx, task, asynq.Queue(queueName))
	if err != nil {
		log.Printf("任务入队失败 [%s]: %v", queueName, err)
	} else {
		log.Printf("任务已入队 ID=%s 队列=[%s] Subject=%s", info.ID, info.Queue, subject)
	}
}
