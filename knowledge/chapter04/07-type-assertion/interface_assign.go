package main

import (
	"fmt"
	"reflect"
)

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

func main() {
	var num1 Number = 1
	var num2 Number2 = &num1
	// 这个表达式断言 num2 是否是 Number1 类型的实例，
	// 如果是，ok 值为 true，然后执行 if 语句块中的代码；
	// 否则 ok 值为 false，不执行 if 语句块中的代码
	if num3, ok := num2.(Number1); ok {
		fmt.Println(num3.Equal(1))        // true
		fmt.Println(reflect.TypeOf(num3)) // *main.Number
	}
}
