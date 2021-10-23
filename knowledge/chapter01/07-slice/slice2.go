package main

import "fmt"

func main() {
	// 定义一个长度为 3 ，容量为 5 的切片
	var numbers = make([]int, 3, 5)

	// len =3, cap = 5, slice = [0 0 0]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向 numbers 切片追加一个元素 11，numbers 的长度为 4，为 [0, 0, 0, 11]，容量还是 5
	numbers = append(numbers, 11)
	// len =4, cap = 5, slice = [0 0 0 11]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 22)
	// len =5, cap = 5, slice = [0 0 0 11 22]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 如果动态增加的元素个数超过了预先设置的容量，那么此时的容量会翻倍，这里其实也就是 5 * 2 = 10
	numbers = append(numbers, 33)
	// len =6, cap = 10, slice = [0 0 0 11 22 33]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 44)
	// len =7, cap = 10, slice = [0 0 0 11 22 33 44]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

}
