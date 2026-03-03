package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// 创建一个新的 Scheduler 实例
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: RedisAddr, DB: RedisDB},
		&asynq.SchedulerOpts{
			Location: time.Local, // 使用本地时区
		},
	)

	// 定义任务：每隔 10 秒执行一次
	// 使用 cron 表达式 "@every 10s"
	task := asynq.NewTask(TypeCronGreeting, nil)

	// 注册任务
	// 注意：为了避免混淆，我们指定这个任务进入 QueueCronTask 队列
	entryID, err := scheduler.Register("@every 10s", task, asynq.Queue(QueueCronTask))
	if err != nil {
		log.Fatalf("无法注册任务: %v", err)
	}
	log.Printf("已注册定时任务 entryID=%q", entryID)

	// 运行 Scheduler
	// Run() 会阻塞并等待信号
	log.Println("正在启动 Scheduler...")
	if err := scheduler.Run(); err != nil {
		log.Fatalf("无法运行 Scheduler: %v", err)
	}
}
