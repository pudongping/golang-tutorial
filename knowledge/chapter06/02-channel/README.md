# 管道

- [channel 的基本定义和使用](./channel.go)
- [channel 有缓冲与无缓冲同步问题](./channel1.go)
- [channel 关闭的特点](./channel2.go)
- [往 goroutine 中写数据和取数据](./channel3.go)
- [定义一个只读只写的 channel](./channel4.go)
- [从只读和只写的 channel 中读写数据](./channel5.go)
- [channel 和 range](./channel_and_range.go)
- [channel 和 select](./channel_and_select.go)
- [select 超时处理](./select_timeout.go)
- [多个 goroutine 抓取数据，并将结果数据以及错误信息保存](./channel_with_err.go)

channel 的基本定义

```go

// 定义一个 channel，Type 为类型
make(chan Type)  // 等价于 make(chan Type, 0)
// 比如定义
c := make(chan int)

// 定义一个有缓冲的
make(chan Type, capacity)

// 将数据写入管道中
channel <- value // 发送 value 到 管道变量 channel 中

// 从 channel 中取值，有以下三种方式
// 第一种
<-channel  // 接收并将其丢弃（从管道中读出数据，但是并没有捕获对应的值）
// 第二种
x := <-channel  // 从 channel 中接收数据，并赋值给 x
// 第三种
x, ok := <-channel  // 功能同上，同时检查通道是否已关闭或者是否为空（ok，表示是否读成功，当 ok 为 true 时，表示通道没有被关闭）


```

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

select 相关：

1. select 只能用于 channel 的操作（写入/读取）
2. select 的 case 是随机的
3. select 要注意避免出现死锁，同时也可以自行实现超时机制
4. select 里面没有类似 switch 里的 fallthrough 的用法
5. select 不能像 switch 一样接函数或其他表达式