/*
匿名函数的定义和使用
*/
package main

import "fmt"

func main() {

	// 将匿名函数赋值给变量
	add := func(a, b int) int {
		return a + b
	}
	// 调用匿名函数
	fmt.Printf("add(1, 2) = %d\n", add(1, 2)) // add(1, 2) = 3

	// 定义时直接调用匿名函数
	func(a, b int) {
		fmt.Println(a + b) // 33
	}(11, 22)

}
