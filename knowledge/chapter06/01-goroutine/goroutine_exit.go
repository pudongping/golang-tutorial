/*
退出当前的 goroutine
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// 用 go 创建承载一个行参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			// 退出当前 goroutine
			// runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	/*
	   // 定义一个匿名函数，且有参数和返回值
	   	go func(a int, b int) bool {
	   		fmt.Println("a =", a, ", b =", b)
	   		return true
	   	}(10, 20)
	*/

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
