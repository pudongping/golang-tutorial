package main

import "fmt"

// 基于 int 设置 Integer 类型
type Integer int

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}

// 是否相等
func (a Integer) Equal(b Integer) bool {
	return a == b
}

// 加法
func (a Integer) Add(b Integer) Integer {
	return a + b
}

// 乘法
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

func main() {
	var x, y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y)) // false

	var a Integer = 1
	var m Math = &a
	fmt.Println(m.Add(1)) // 2
}
