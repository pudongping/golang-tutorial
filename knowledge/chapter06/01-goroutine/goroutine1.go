package main

import (
	"fmt"
	"sync"
)

func main() {

	// 声明一个协程管理器
	var wg sync.WaitGroup
	wg.Add(5) // 添加 5 个协程
	for i := 0; i < 5; i++ {
		go Run(&wg, i)
	}

	// 让协程执行等待
	wg.Wait()  // 阻塞，会等待所有的 goroutine 执行完毕
}

func Run(wg *sync.WaitGroup, i int) {
	fmt.Println("hello alex", i)
	wg.Done() // 关闭掉协程，让计数器减 1
}
