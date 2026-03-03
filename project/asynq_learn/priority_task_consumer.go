package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// HandlePriorityEmailTask 处理优先级任务
func HandlePriorityEmailTask(ctx context.Context, t *asynq.Task) error {
	var payload map[string]string
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal 失败: %v: %w", err, asynq.SkipRetry)
	}

	// 打印处理日志，注意观察输出顺序
	// 实际上这里应该有业务逻辑，比如发送邮件
	log.Printf("正在处理任务: [队列=%s] [用户=%s] [主题=%s]",
		payload["queue"], payload["user_id"], payload["subject"])

	// 模拟处理耗时，以便观察并发时的优先级效果
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	return nil
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB},
		asynq.Config{
			// 并发数设置为 1，这样可以更清晰地看到处理顺序（严格串行）
			// 如果并发数较高，可能会因为并发执行而显得顺序没那么严格，但统计上高优先级任务会被更多地获取
			Concurrency: 1,
			// 核心配置：队列优先级权重
			// 数字越大，优先级越高（被处理的概率越大）
			// 这里配置 6:3:1 的比例
			Queues: map[string]int{
				QueuePriorityPlatinum: 6, // 铂金会员：最高权重
				QueuePriorityDiamond:  3, // 钻石会员：中等权重
				QueuePriorityNormal:   1, // 普通会员：最低权重
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Printf("任务处理失败: %v", err)
			}),
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailSend, HandlePriorityEmailTask)

	log.Println("正在启动 Asynq 优先级队列消费者...")
	log.Println("权重配置 -> 铂金:6, 钻石:3, 普通:1")

	if err := srv.Run(mux); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
