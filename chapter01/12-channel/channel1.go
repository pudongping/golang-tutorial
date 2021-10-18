/*
channel 有缓冲与无缓冲同步问题
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3) // 带有缓冲的 channel

	// 打印 c 的长度和容量
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("子 go 程结束")

		// 特点1：当 channel 已经满了，再向里面写数据时，就会阻塞
		// 这里的 c 的空间定义时，容量只有 3，但是此时却往 c 里面写了 4 次数据
		for i := 0; i < 4; i++ {
			c <- i // 将 i 以管道的形式丢给变量 c
			fmt.Println("子 go 程正在运行，发送的元素 =", i, " len(c)=", len(c), ", cap(c)=", cap(c))
		}

	}()

	time.Sleep(2 * time.Second)

	// 特点2：当 channel 已经为空，从里面取数据也会阻塞
	for i := 0; i < 4; i++ {
		num := <-c // 从 c 中接收数据，并赋值给 num
		fmt.Println("num =", num)
	}

	fmt.Println("main 结束")

}

/*
len(c) =  0 , cap(c) 3
子 go 程正在运行，发送的元素 = 0  len(c)= 1 , cap(c)= 3
子 go 程正在运行，发送的元素 = 1  len(c)= 2 , cap(c)= 3
子 go 程正在运行，发送的元素 = 2  len(c)= 3 , cap(c)= 3
num = 0
num = 1
num = 2
num = 3
main 结束
*/
