/*
协程的创建
*/
package main

import (
	"fmt"
	"time"
)

// 子 goroutine
// 主 goroutine 停止之后，子 goroutine 也会对应的停止
func newTask() {
	i := 0
	// 死循环
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {

	// 创建一个 go 程，去执行 newTask() 流程
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
