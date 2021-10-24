/*
按值传参
*/
package main

import "fmt"

func add(a, b int) int {
	// 这里的变量 a 和变量 b，其实是由实参变量 x 和变量 y，拷贝出的一个副本进行赋值的
	a *= 2
	b *= 3
	return a + b
}

func main() {

	x, y := 1, 2
	z := add(x, y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)       // add(1, 2) = 8
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z) // x = 1, y = 2, z = 8

}
