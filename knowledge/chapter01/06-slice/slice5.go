package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3]
	fmt.Println("slice1 =", slice1) // slice1 = [1 2 3 4 5]
	fmt.Println("slice2 =", slice2) // slice2 = [2 3]
	slice2[1] = 6
	// 此时改变了 slice2 但是 slice1 也跟着改变了，因为是指针指向了同一个数组
	fmt.Println("slice1 =", slice1) // slice1 = [1 2 6 4 5]
	fmt.Println("slice2 =", slice2) // slice2 = [2 6]

	fmt.Println("===============优美的分割线===================")
	slice3 := make([]int, 4)
	fmt.Println("slice3 =", slice3) // slice3 = [0 0 0 0]
	slice4 := slice3[1:3]
	fmt.Println("slice4 =", slice4) // slice4 = [0 0]
	slice3 = append(slice3, 0)      // 此时切片的容量已经超过了 4 ，因此程序底层会对切片进行了扩容操作，会重新分配内存空间，因此就不会发生数据共享问题
	fmt.Println("slice3 =", slice3) // slice3 = [0 0 0 0 0]
	slice3[1] = 2
	fmt.Println("slice3 =", slice3) // slice3 = [0 2 0 0 0]
	slice4[1] = 6
	fmt.Println("slice4 =", slice4) // slice4 = [0 6]
	fmt.Println("=============")
	fmt.Println("slice3 =", slice3) // slice3 = [0 2 0 0 0]
	fmt.Println("slice4 =", slice4) // slice4 = [0 6]

}
