package main

import "fmt"

var v1 int  // 整型
var v2 string  // 字符串
var v3 bool  // 布尔型
var v4 [10]int  // 数组，数组元素类型为整型
var v5 struct{  // 结构体，成员变量 f 的类型为 64 位浮点型
	f float64
}
var v6 *int  // 指针，指向整型
var v7 map[string]int  // map （字典），key 为字符串类型，value 为整型
var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型

func main()  {

	fmt.Println("v1 :", v1)  // v1 : 0
	fmt.Println("v2 :", v2)  // v2 :
	fmt.Println("v3 :", v3)  // v3 : false
	fmt.Println("v4 :", v4)  // v4 : [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("v5 :", v5)  // v5 : {0}
	fmt.Println("v6 :", v6)  // v6 : <nil>
	fmt.Println("v7 :", v7)  // v7 : map[]
	fmt.Println("v8 :", v8)  // v8 : <nil>

	// 变量多重赋值
	i := 11
	j := 22
	fmt.Println("i =", i)  // i = 11
	fmt.Println("j =", j)  // j = 22
	i, j = j, i  // 交换变量 i 和 j 的值
	fmt.Println("i =", i)  // i = 22
	fmt.Println("j =", j)  // j = 11

}