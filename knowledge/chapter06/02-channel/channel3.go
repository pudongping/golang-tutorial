package main

import (
	"fmt"
)

func main() {

	c1 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c1)
	}

}
