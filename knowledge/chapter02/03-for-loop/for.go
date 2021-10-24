package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	// 1 + 2 + 3 + 4 + ... + 98 + 99 + 100 =  5050
	fmt.Println("1 + 2 + 3 + 4 + ... + 98 + 99 + 100 = ", sum)

	fmt.Println("================= 优美的分割线 =================")

	// 无限循环
	sum1 := 0
	i := 0
	for {
		i++
		if i > 100 {
			break // 可以通过 break 语句来中断无限循环
		}
		sum1 += i
	}
	// 1 + 2 + 3 + 4 + ... + 99 + 100 =  5050
	fmt.Println("1 + 2 + 3 + 4 + ... + 99 + 100 = ", sum1)

	fmt.Println("================= 优美的分割线 =================")

	// 多重赋值
	// 快速实现数组/切片内首尾元素的交换
	arr := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	// new arr is  [6 5 4 3 2 1]
	fmt.Println("new arr is ", arr)

	fmt.Println("================= 优美的分割线 =================")

	// for-range 结构
	for k, v := range arr {
		fmt.Println("k =", k, "v =", v)
	}
	/*
		k = 0 v = 6
		k = 1 v = 5
		k = 2 v = 4
		k = 3 v = 3
		k = 4 v = 2
		k = 5 v = 1
	*/

	fmt.Println("================= 优美的分割线 =================")

	// 循环过程中，要忽略索引或者键
	for _, v := range arr {
		fmt.Println("v1 =", v)
	}
	/*
		v1 = 6
		v1 = 5
		v1 = 4
		v1 = 3
		v1 = 2
		v1 = 1
	*/

	fmt.Println("================= 优美的分割线 =================")

	// 要忽略元素值
	for k := range arr {
		fmt.Println("k1 =", k)
	}
	/*
		k1 = 0
		k1 = 1
		k1 = 2
		k1 = 3
		k1 = 4
		k1 = 5
	*/

	fmt.Println("================= 优美的分割线 =================")

	// 基于条件判断进行循环
	sum2 := 0
	i2 := 0
	for i2 < 4 {
		i2++
		sum2 += i2
	}
	fmt.Println("sum2 =", sum2) // sum2 = 10

	fmt.Println("================= 优美的分割线 =================")

ABCLoop:
	for j3 := 0; j3 < 5; j3++ {
		for i3 := 0; i3 < 10; i3++ {
			if i3 > 5 {
				break ABCLoop // break 语句终止的是 ABCLoop 标签处的外层循环
			}
			fmt.Println("i3 =", i3)
		}
	}
	/*
		i3 = 0
		i3 = 1
		i3 = 2
		i3 = 3
		i3 = 4
		i3 = 5
	*/

}
