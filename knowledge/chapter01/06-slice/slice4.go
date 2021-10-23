package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// len =10, cap = 10, slice = [1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)
	slice1 = slice1[:len(slice1)-4] // 删除 slice1 尾部的 4 个元素
	// len =6, cap = 10, slice = [1 2 3 4 5 6]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)
	slice1 = slice1[5:] // 删除 slice1 头部 5 个元素
	// len =1, cap = 5, slice = [6]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("a =", slice1[:0], "b =", slice1[3:]) // a = [] b = [4 5 6 7 8 9 10]
	slice2 := append(slice1[:0], slice1[3:]...)       // 删除开头三个元素
	// len =7, cap = 10, slice = [4 5 6 7 8 9 10]
	fmt.Printf("len =%d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)

}
