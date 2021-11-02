/*
基本数据类型，通过 variable.(type) 获取变量对应的类型值
*/
package main

import "fmt"

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case bool:
			fmt.Println(arg, "is a bool value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	abc := "alex"
	myPrintf(abc) // "alex" is a string value.
}
