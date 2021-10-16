/*
通过反射获取结构体中的属性以及方法
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func main() {

	user := User{1, "Alex", 18}

	DoFiledAndMethod(user)

}

func DoFiledAndMethod(input interface{}) {

	// 获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name()) // inputType is: User

	// 获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue) // inputValue is: {1 Alex 18}

	// 通过 type 获取里面的字段
	// 1. 获取 interface 的 reflect.Type ，通过 Type 得到 NumField，进行遍历
	// 2. 得到每个 field，数据类型
	// 3. 通过 field 有一个 Interface() 方法得到对应的 value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		// Id: int = 1
		// Name: string = Alex
		// Age: int = 18

	}

	// 通过 type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s => %v\n", m.Name, m.Type)
		// Call => func(main.User)
	}

}
