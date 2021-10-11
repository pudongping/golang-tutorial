/**
变量
 */
package main

import "fmt"

/**
四种变量的声明方式
 */

// 声明全局变量：方式一、方式二、方式三是可以的
var gV1 string
var gV2 string = "gV2 ===> alex"
var gV3 = "gV3 ===> alex"
// := 只能够用在函数体内来声明
//gV4 := "alex"

func main()  {

	// 方式一：声明一个变量，默认的值是 0
	var v1 int
	fmt.Println("v1 = ", v1)  // v1 =  0
	fmt.Printf("type of v1 = %T\n", v1)  // type of v1 = int

	// 方式二：声明一个变量，初始化一个值
	var v2 int = 100
	fmt.Println("v2 = ", v2)  // v2 =  100
	fmt.Printf("type of v2 = %T\n", v2)  // type of v2 = int
	var v3 string = "abcd"
	fmt.Printf("v3 = %s, type of v3 = %T\n", v3, v3)  // v3 = abcd, type of v3 = string

	// 方式三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var v4 = 100
	fmt.Println("v4 = ", v4)  // v4 =  100
	fmt.Printf("type of v4 = %T\n", v4)  // type of v4 = int
	var v5 = "abcd"
	fmt.Printf("v5 = %s, type of v5 = %T\n", v5, v5)  // v5 = abcd, type of v5 = string

	// 方式四：（常用的方式）省去 var 关键字，直接自动匹配
	v6 := 3.14
	fmt.Println("v6 = ", v6)  // v6 =  3.14
	fmt.Printf("type of v6 = %T\n", v6)  // type of v6 = float64

	// 打印全局变量时
	fmt.Printf("global variable gV1 = [ %s ], type of gV1 = %T\n", gV1, gV1)  // global variable gV1 = [  ], type of gV1 = string
	fmt.Printf("global variable gV2 = [ %s ], type of gV2 = %T\n", gV2, gV2)  // global variable gV2 = [ gV2 ===> alex ], type of gV2 = string
	fmt.Printf("global variable gV3 = [ %s ], type of gV3 = %T\n", gV3, gV3)  // global variable gV3 = [ gV3 ===> alex ], type of gV3 = string

	// 声明多个变量
	var v7, v8 int = 100, 200  // 声明同一类型的变量时
	fmt.Printf("v7 = %d, v8 = %d\n", v7, v8)  // v7 = 100, v8 = 200
	var v9, v10 = 100, "alex"  // 声明不同类型的变量时
	fmt.Println("v9 = ", v9, ", v10 = ", v10)  // v9 =  100 , v10 =  alex
	// 多行的多变量声明
	var (
		v11 int = 100
		v12 bool = true
	)
	fmt.Println("v11 =", v11, ", v12 =", v12)  // v11 = 100 , v12 = true

}