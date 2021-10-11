/*
定义固定长度的数组
 */
package main

import "fmt"

func printArray(myArray [4]int)  {

	// 传过来的参数为值拷贝过程
	for index, value := range myArray {
		fmt.Println("index =", index, "value =", value)
	}

	myArray[0] = 111  // 修改数组中的值

}

func main()  {

	// 固定长度为 10 的数组
	var myArray1 [10]int  // 数组中所有的值都为 0

	// 遍历数组方式一
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// 前 4 个元素是有值的，后 6 个元素默认值为 0
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	// 遍历数组方式二
	for index, value := range myArray2 {
		fmt.Println("index =", index, "value =", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	// 打印数组
	printArray(myArray3)
	fmt.Println(" --------- ")
	// 因为在 printArray 方法中已经尝试修改了传参，那么这里看是否会已经修改了传递的参数
	// 通过实验可知，根本就没有改变原始 myArray3 数组中的值
	for index, value := range myArray3 {
		fmt.Println("myArray3 index =", index, "myArray3 value =", value)
	}

}


/*
0
0
0
0
0
0
0
0
0
0
index = 0 value = 1
index = 1 value = 2
index = 2 value = 3
index = 3 value = 4
index = 4 value = 0
index = 5 value = 0
index = 6 value = 0
index = 7 value = 0
index = 8 value = 0
index = 9 value = 0
myArray1 types = [10]int
myArray2 types = [10]int
myArray3 types = [4]int
index = 0 value = 11
index = 1 value = 22
index = 2 value = 33
index = 3 value = 44
 ---------
myArray3 index = 0 myArray3 value = 11
myArray3 index = 1 myArray3 value = 22
myArray3 index = 2 myArray3 value = 33
myArray3 index = 3 myArray3 value = 44
 */