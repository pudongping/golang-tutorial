# 路由

这里演示的是：将各种消息发送给交换器，接收者根据监听不同的路由来接收消息。
就好比在日志服务中，有各种各样类型的日志，但是 A 消费者只需要 warning 级别的日志、B 消费者只需要 danger 级别的日志一样。

- `fanout` 交换器并不是很灵活，它只能进行无脑广播。（会将所有消息广播给所有的消费者）
- `direct` 交换器 —— 消息进入其 `binding key` 与消息的 `routing key` 完全匹配的队列。

使用示例

启动一个终端 A

```bash
# 比如这里，只想接收 info、warning 级别的消息
go run subscribe_direct.go info warning
```

再启动一个终端 B

```bash
# 这里只想接收 error 级别的消息
go run subscribe_direct.go error
```

再启动一个终端 C，作为生产者

```bash
# 投递一条 notice 级别的消息（这条消息将被丢弃掉，因为没有消费者监听这个 notice 路由）
go run publish_direct.go notice "hello world" 

# 投递一条 info 级别的消息，这条消息将被 info 路由收到，也就是会被终端 A 打印出来，终端 B 不会有任何输出
go run publish_direct.go info "hello world"

# 投递一条 error 级别的消息，这条消息将被 error 路由收到，也就是会被终端 B 打印出来
go run publish_direct.go error "hello world"
```