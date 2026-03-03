package main

const (
	RedisAddr = "localhost:6379"
	RedisDB   = 4
)

// 队列名称
const (
	QueueAsyncTaskEmailDefault = "async_task:email:default"
)

// 任务类型
const (
	TypeEmailSend = "email:send"
)
