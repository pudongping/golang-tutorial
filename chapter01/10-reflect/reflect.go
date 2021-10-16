/*
通过反射获取变量的类型以及数据值
 */
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))   // type:  float64
	fmt.Println("value: ", reflect.ValueOf(arg)) // value:  1.2345
}

func main() {

	var num float64 = 1.2345

	reflectNum(num)

}
