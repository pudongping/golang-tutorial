package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c

	go SetChan(writec)
	GetChan(readc)

}

func SetChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		writec <- i
	}
}

func GetChan(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("从 GetChan 函数中取的值 %d \n", <-readc)
	}
}
