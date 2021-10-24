/*
任意类型的变长参数（泛型）
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
	myPrintf(1, "1", [1]int{1}, true, false, "false")
	/*
		1 is an int value.
		"1" is a string value.
		[1] is an array type.
		true is an unknown type.
		false is an unknown type.
		"false" is a string value.
	*/

}
