package main

const (
	RedisAddr = "localhost:6379"
	RedisDB   = 4
)

// 队列名称
const (
	QueueAsyncTaskEmailDefault = "async_task:email:default"
	QueueDelayTaskEmailDefault = "delay_task:email:default"
	// 优先级队列
	QueuePriorityPlatinum = "priority:email:platinum" // 铂金
	QueuePriorityDiamond  = "priority:email:diamond"  // 钻石
	QueuePriorityNormal   = "priority:email:normal"   // 普通
	// 定时任务队列
	QueueCronTask = "cron:greeting"
)

// 任务类型
const (
	TypeEmailSend    = "email:send"
	TypeCronGreeting = "cron:greeting"
)
