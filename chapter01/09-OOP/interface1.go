/*
interface{} 空接口万能类型与类型断言机制
 */
package main

import "fmt"

func myFunc(arg interface{})  {
	fmt.Println("myFunc is called ...")
	fmt.Println("arg =", arg)

	// interface{} 该如何区分此时引用的底层数据类型到底是什么？

	// 给 interface{} 提供 "类型断言" 的机制
	value, ok := arg.(string)  // 判断 arg 参数是否为 string 类型
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("value type is %T\n", value)  // 打印此时 arg 对应的 value 值出来
	}

}

type Books struct {
	auth string
}

func main()  {
	book := Books{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)

}

/*

myFunc is called ...
arg = {Golang}
arg is not string type
myFunc is called ...
arg = 100
arg is not string type
myFunc is called ...
arg = abc
arg is string type, value = abc
value type is string
myFunc is called ...
arg = 3.14
arg is not string type

 */