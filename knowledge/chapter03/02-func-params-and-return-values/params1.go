/*
引用传参
*/
package main

import "fmt"

func addPlus(a, b *int) int {
	*a *= 2
	*b *= 3

	return *a + *b
}

func main() {

	x, y := 1, 2
	// 此时传递给函数的参数是一个指针，而指针代表的是实参的内存地址，
	// 修改指针引用的值即修改变量内存地址中存储的值，所以实参的值也会被修改（
	// 这种情况下，传递的是变量地址值的拷贝，所以从本质上来说还是按值传参）
	z := addPlus(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)       // add(2, 6) = 8
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z) // x = 2, y = 6, z = 8

}
