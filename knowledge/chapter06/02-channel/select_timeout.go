package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)

	// 当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞，
	// 但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间
	go handlerTimeout(timeout, 2)

	select {
	case msg1 := <-c1:
		fmt.Println("c1 Received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 Received: ", msg2)
	case <-timeout:
		fmt.Println("timeout exit.")
	}

}

func handlerTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}
