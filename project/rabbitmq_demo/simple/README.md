
一个简单的生产者和消费者示例代码

先运行生产者

```bash
go test -run TestProducerMessage
```

再运行消费者

```bash
go test -run TestConsumerMessage
```

如果需要检查队列，可以在 rabbitmq docker 容器中，使用 `rabbitmqctl list_queues` 命令来查看。
