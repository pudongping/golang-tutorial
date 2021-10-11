/**
声明切片的 4 种方式
 */
package main

import "fmt"

func main()  {

	// 声明切片的方式一：
	// 声明 slice1 是一个切片，并且初始化，默认值是 1, 2, 3 。长度 len 是 3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)  // len = 3, slice = [1 2 3]

	// 声明切片的方式二：
	// 声明 slice2 是一个切片，但是并没有给 slice2 分配空间
	var slice2 []int
	slice2 = make([]int, 3)  // 给数组开辟 3 个空间，默认值是 0
	fmt.Printf("slice2 = %v\n", slice2)  // slice2 = [0 0 0]
	slice2[0] = 100
	fmt.Printf("slice2 = %v\n", slice2)  // slice2 = [100 0 0]

	// 声明切片的方式三：
	// 声明 slice3 是一个切片，同时给 slice3 分配 3 个空间，初始值是 0
	var slice3 []int = make([]int, 3)
	fmt.Printf("slice3 = %v\n", slice3)  // slice3 = [0 0 0]

	// 声明切片的方式四：
	// 通过 := 推导出 slice4 是一个切片
	// 声明 slice4 是一个切片，同时给 slice4 分配 3 个空间，初始化值是 0
	slice4 := make([]int, 3)
	fmt.Printf("slice4 = %v\n", slice4)  // slice4 = [0 0 0]

	// 判断一个 slice 是否为 0
	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5 是一个空切片")
	} else {
		fmt.Println("slice5 是有空间的")
	}

}