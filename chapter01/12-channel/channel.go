/*
channel 的基本定义和使用
*/
package main

import "fmt"

func main() {

	// 定义一个 channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行……")

		c <- 888 // 将 888 发送给 c 变量

	}()

	fmt.Println("haha")
	num := <-c // 从变量 c 中接收数据，并赋值给 num

	fmt.Println("num =", num)
	fmt.Println("main goroutine 结束……")

}

/*
haha
goroutine 正在运行……
goroutine 结束
num = 888
main goroutine 结束……
*/
