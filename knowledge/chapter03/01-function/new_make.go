/*
在 Go 语言中，引用类型包括切片（slice）、字典（map）和管道（channel），其它都是值类型。
*/
package main

import "fmt"

func main() {

	// new 函数作用于值类型，仅分配内存空间，返回的是指针
	v1 := new(int)    // 返回 int 类型指针，相当于 var v1 *int
	v2 := new(string) // 返回 string 类型指针
	v3 := new([3]int) // 返回数组类型指针，数组长度是 3

	type Student struct {
		id    int
		name  string
		grade string
	}
	v4 := new(Student) // 返回对象类型指针

	println("v1 =", v1) // v1 = 0xc000056720
	println("v2 =", v2) // v2 = 0xc000056740
	println("v3 =", v3) // v3 = 0xc000056728
	println("v4 =", v4) // v4 = 0xc000056750

	// make 函数作用于引用类型，除了分配内存空间，还会对对应类型进行初始化，返回的是初始值
	// 引用类型包括切片（slice）、字典（map）、和管道（channel），其它都是值类型

	v5 := make([]int, 3)
	v6 := make(map[string]int, 2)
	println("v5 =", v5, "len is", len(v5)) // v5 = [3/3]0xc000014090 len is 3
	println("v6 =", v6, "len is", len(v6)) // v6 = 0xc00007a180 len is 0
	fmt.Println("v5 =", v5)                // v5 = [0 0 0]
	fmt.Println("v6 =", v6)                // v6 = map[]

	for k, v := range v5 {
		println("k =", k, "v =", v)
	}
	/*
		k = 0 v = 0
		k = 1 v = 0
		k = 2 v = 0
	*/

	println("============== 优美的分割线 =============")

	v6["hello"] = 100
	for k1, v1 := range v6 {
		println("k1 =", k1, "v1 =", v1)
	}
	/*
		k1 = hello v1 = 100
	*/

}
