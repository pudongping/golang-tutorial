package main

import "fmt"

func main()  {

	// 空接口指向基本类型
	var v1 interface{} = 1  // 将 int 类型赋值给 interface{}
	var v2 interface{} = "Alex"  // 将 string 类型赋值给 interface{}
	var v3 interface{} = true  // 将 bool 类型赋值给 interface{}

	// 空接口指向复合类型
	var v4 interface{} = &v2  // 将指针类型赋值给 interface{}
	var v5 interface{} = []int{1, 2, 3}  // 将切片类型赋值给 interface{}
	var v6 interface{} = struct {  // 将结构体类型赋值给 interface{}
		id int
		name string
	}{123, "Alex"}

	fmt.Println("v1 =", v1)  // v1 = 1
	fmt.Println("v2 =", v2)  // v2 = Alex
	fmt.Println("v3 =", v3)  // v3 = true
	fmt.Println("v4 =", v4)  // v4 = 0xc000108050
	fmt.Println("v5 =", v5)  // v5 = [1 2 3]
	fmt.Println("v6 =", v6)  // v6 = {123 Alex}

}
