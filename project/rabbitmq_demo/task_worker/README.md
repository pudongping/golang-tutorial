# 工作队列、任务队列

演示步骤

需要打开 3 个终端

启动 2 个消费者

第一个消费者

```bash
go run worker.go
```

第二个消费者

```bash
go run worker.go
```

> 以上两个消费者都将从队列中获取消息

启动生产者

```bash
# 不断的往队列里面塞数据，每次增加一个点用来模拟消费者执行耗时
go run task.go nice

.
.
.

go run task.go nice...
go run task.go nice....
go run task.go nice......
go run task.go nice.......
go run task.go nice........
```

当我们启动多个消息消费者时，默认情况下，RabbitMQ 将按顺序将每个消息发送给下一个消费者。
平均而言，每个消费者都会收到**相同数量**的消息。这种分发消息的方式叫做**轮询**。

## 持久化设置

生产者和消费者的队列都需要设置为**持久化**

```go
q, err := ch.QueueDeclare(
    taskQueueName, // 队列名称
    true,          // 持久化（如果不设置为 true，那么当 RabbitMQ 服务器停止运行或者崩溃时，消息就会丢失）
    false,         // 自动删除
    false,         // 排他性
    false,         // 等待服务器确认
    nil,           // 参数
)
```

并且生产者中，需要将消息标记为持久的

```go
err = ch.PublishWithContext(
    ctx,
    "",     // 交换机名称
    q.Name, // 队列名称
    false,  // 必需的
    false,  // 立即发布
    amqp.Publishing{
        ContentType:  "text/plain",
        DeliveryMode: amqp.Persistent, // 持久（交付模式：瞬态/持久）=> 将消息标记为持久的（在队列中标记为“持久化”还不行，还一定需要在发送消息的时候标记为“持久”）
        Body:         []byte(body),
    },
)
```

### 有关消息持久性的说明

将消息标记为持久性并不能完全保证消息不会丢失。尽管它告诉 RabbitMQ 将消息保存到磁盘上，但是 RabbitMQ 接受了一条消息并且还没有保存它时，仍然有一个很短的时间窗口。而且，RabbitMQ 并不是对每个消息都执行 `fsync(2)`——它可能只是保存到缓存中，而不是真正写入磁盘。
持久性保证不是很强，但是对于我们的简单任务队列来说已经足够了。如果您需要更强有力的担保，那么您可以使用 [publisher confirms](https://www.rabbitmq.com/confirms.html)。

## 公平分发