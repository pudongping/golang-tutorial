/*
channel 和 select
select 具备多路 channel 的监控状态功能
*/
package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			// 如果 c 可写，则该 case 就会进来
			x = y
			y = x + y
		case <-quit:
			// 如果 quit 可读，则该 case 就会进来
			fmt.Println("quit")
			return
		}
	}

}

func main() {

	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {

		for i := 0; i < 10; i++ {
			// 这里触发 c 的读操作
			fmt.Println(<-c)
		}

		// 这里触发 quit 写的操作
		quit <- 0
	}()

	// main go
	fibonacii(c, quit)
}
