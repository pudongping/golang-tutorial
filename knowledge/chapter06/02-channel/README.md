# 管道

- [channel 的基本定义和使用](./channel.go)
- [channel 有缓冲与无缓冲同步问题](./channel1.go)
- [channel 关闭的特点](./channel2.go)
- [channel 和 range](./channel_and_range.go)
- [channel 和 select](./channel_and_select.go)
- [往 goroutine 中写数据和取数据](./channel3.go)
- [定义一个只读只写的 channel](./channel4.go)
- [从只读和只写的 channel 中读写数据](./channel5.go)

单流程下一个 go 只能监控一个 channel 的状态，select 可以完成监控多个 channel 的状态

```go

select {
case <- chan1:
	// 如果 chan1 成功读到数据，则进行该 case 处理语句
case chan2 <- 1:
	// 如果成功向 chan2 写入数据，则进行该 case 处理语句
default:
	// 如果上面都没有成功，则进入 default 处理流程
}

```