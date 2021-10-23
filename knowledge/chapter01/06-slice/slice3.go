package main

import "fmt"

func main() {

	// 长度为 3，容量也为 3， [1, 2, 3]
	s := []int{1, 2, 3}

	// 截取 [0, 2) 长度的元素，亦即截取后的结果为 [1, 2]
	s1 := s[0:2]
	fmt.Println(s1) // [1 2]

	// s1 虽然是通过截取的 s ，但是他们两者指向的内存地址都是同一个，因此修改截取后的切片，那么原始切片的值也会被改变
	s1[0] = 100
	fmt.Println(s)  // [100 2 3]
	fmt.Println(s1) // [100 2]

	// copy 函数可以将底层数组的 slice 一起进行深拷贝
	s2 := make([]int, 3) // s2 = [0 0 0]
	fmt.Println(s2)      // [0 0 0]

	// 内容复制
	// 将 s 中的值，依次拷贝到 s2 中
	copy(s2, s)
	fmt.Println(s2) // [100 2 3]

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	// 复制 slice1 到 slice2
	copy(slice2, slice1)            // 只会复制 slice1 的前 3 个元素到 slice2 中
	fmt.Println("slice2 =", slice2) // slice2 = [1 2 3]

	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{5, 4, 3}
	// 复制 slice4 到 slice3
	copy(slice3, slice4)            // 只会复制 slice4 的 3 个元素到 slice3 的前 3 个位置
	fmt.Println("slice3 =", slice3) // slice3 = [5 4 3 4 5]

}
