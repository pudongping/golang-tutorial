/*
channel 关闭的特点
1. channel 不像文件一样需要经常去关闭，只有当你确定没有任何发送数据了，或者你想显式的结束 range 循环之类的，才去关闭 channel
2. 关闭 channel 后，无法向 channel 再发送数据 （引发 panic 错误后导致接收立即返回零值）
3. 关闭 channel 后，可以继续从 channel 接收数据
4. 对于 nil channel，无论收发都会被阻塞
5. 对于已经关闭的 channel，再次关闭会引发 panic 错误
*/
package main

import "fmt"

func main() {

	c := make(chan int) // 创建一个无缓冲的 channel

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 当我们向通道中发送完数据后，我们可以通过 close 来关闭通道 channel
		// close 可以关闭一个 channel
		// 通道的关闭操作应该由发送方来完成，接收方不应该关闭通道
		close(c)
	}()

	for {
		// 变量 ok 如果为 true 表示 channel 没有关闭，如果为 false 表示 channel 已经关闭
		// 这里的 data 和 ok 其实在 if 判断语句中属于局部变量
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished ……")
}

/*
0
1
2
3
4
Main Finished ……
*/
