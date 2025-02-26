## 场景1: 消息正常消费

先执行生产者

```bash
go run producer.go
```

再执行消费者

```bash
go run consumer.go
```

## 场景2: 消息过期测试

单独运行生产者发送消息，我们会发现，所有进入队列的消息 10 秒后自动删除

验证方法：

```bash
# 查看队列状态
rabbitmqctl list_queues name messages_ready messages_unacknowledged

# 查看消费者连接
rabbitmqctl list_connections
```

运行生产者后观察消息数的变化。

---

生产节奏影响：

- 生产者每 1 秒发 1 条消息
- 第 10 条消息在第 9 秒发送
- 所有消息将在第 19 秒前全部过期（第 1 条在第 10 秒过期，第 10 条在第 19 秒过期）

等待 20 秒后，再启动消费者 `go run consumer.go`，可以看到消费者无任何信息输出。