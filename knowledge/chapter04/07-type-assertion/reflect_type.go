/*
使用 reflect 包提供的 TypeOf 函数动态断言类型
*/
package main

import (
	"fmt"
	"reflect"
)

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	abc := 111
	myPrintf(abc) // 111 is an int value.
}
