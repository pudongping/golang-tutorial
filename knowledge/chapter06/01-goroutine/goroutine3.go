package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "cat"
			time.Sleep(time.Microsecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "dog"
			time.Sleep(time.Microsecond * 500)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}

}
