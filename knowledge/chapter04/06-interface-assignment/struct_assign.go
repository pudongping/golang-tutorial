package main

import "fmt"

type Integer int

// 加法运算
func (a *Integer) Add(b Integer) {
	*a = (*a) + b
}

// 乘法运算
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

func main() {
	var a Integer = 1
	var m Math = &a
	m.Add(2)
	// 1 + 2 = 3
	fmt.Printf("1 + 2 = %d\n", a)
}
