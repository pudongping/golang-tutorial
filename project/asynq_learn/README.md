# asynq 包

通过 [hibiken/asynq](https://github.com/hibiken/asynq) 包实现异步任务队列。


## 1. 异步队列

- [异步队列生产者](async_task_producer.go)
- [异步队列消费者](async_task_consumer.go)

启动方式

```shell
# 启动消费者
go run async_task_consumer.go config.go

# 启动生产者
go run async_task_producer.go config.go 
```