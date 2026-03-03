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

## 2. 延迟队列

- [延迟队列生产者](delay_task_producer.go)
- [延迟队列消费者](delay_task_consumer.go)

演示延迟 30 秒发送邮件。

启动方式

```shell
# 启动延迟任务消费者
go run delay_task_consumer.go config.go

# 启动延迟任务生产者
go run delay_task_producer.go config.go
```

## 3. 优先级队列（加权队列）

- [优先级队列生产者](priority_task_producer.go)
- [优先级队列消费者](priority_task_consumer.go)

演示不同优先级的任务处理顺序。权重配置如下：
- **铂金会员 (Platinum)**: 权重 6 (最高)
- **钻石会员 (Diamond)**: 权重 3 (中等)
- **普通会员 (Normal)**: 权重 1 (最低)

这意味着在大量任务并发时，铂金会员的任务被处理的概率是普通会员的 6 倍。

启动方式

```shell
# 启动优先级任务消费者
# 为了更清晰地观察优先级效果，消费者设置了 Concurrency: 1 (串行处理)
go run priority_task_consumer.go config.go

# 启动优先级任务生产者 (批量发送混合任务)
go run priority_task_producer.go config.go
```

## 4. 定时任务 (Periodic Task)

- [定时任务调度器](periodic_task_scheduler.go)
- [定时任务消费者](periodic_task_consumer.go)

演示每隔 10 秒打印一句问候语。

启动方式

```shell
# 启动定时任务消费者（处理实际逻辑）
go run periodic_task_consumer.go config.go

# 启动调度器（负责定时触发任务）
go run periodic_task_scheduler.go config.go
```
