package main

import "fmt"

func main()  {

	var a string
	// pair<static_type:string, value:"Alex">
	a = "Alex"

	// pair<type:string, value:"Alex">
	var allType interface{}
	allType = a

	// 不管怎么传递，变量里面的数据类型没有发生改变
	value, _ := allType.(string)
	fmt.Println(value)  // Alex

}
