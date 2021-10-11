/*
声明动态数组，切片的方式
 */
package main

import "fmt"

func printSliceArray(myArr []int)  {
	// 切片传参时，是引用传递

	// 当只关心数组中的值，不关心数组的索引时，我们可以使用匿名变量的方式
	// _ 表示匿名的变量
	for _, value := range myArr {
		fmt.Println("value =", value)
	}

	myArr[0] = 11

}

func main()  {

	myArray := []int{1, 2, 3, 4}  // 声明一个动态的数组，传递多少值进去，则该数组的长度就有多长
	fmt.Printf("myArray type is %T\n", myArray)  // myArray type is []int

	printSliceArray(myArray)
	fmt.Println(" ============ ")

	for _, value := range myArray {
		fmt.Println("value =", value)
	}

}

/*
myArray type is []int
value = 1
value = 2
value = 3
value = 4
 ============
value = 11
value = 2
value = 3
value = 4
 */
