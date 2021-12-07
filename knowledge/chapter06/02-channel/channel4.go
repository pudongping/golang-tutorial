package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5)

	// 定义一个只能读的 channel
	var readc <-chan int = c1
	// 定义一个只能写的 channel
	var writec chan<- int = c1

	writec <- 1
	fmt.Println(<-readc)

}
