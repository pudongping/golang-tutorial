package memory_alignment

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int8
	b int8
	c int32
	d string
	e string
}

type B struct {
	a int8
	e string
	c int32
	b int8
	d string
}

type C struct {
	d string
	e string
	c int32
	a int8
	b int8
}

func Run() {
	var a A
	var b B
	var c C
	fmt.Printf("a size: %v \n", unsafe.Sizeof(a))
	fmt.Printf("b size: %v \n", unsafe.Sizeof(b))
	fmt.Printf("c size: %v \n", unsafe.Sizeof(c))
	// a size: 40
	// b size: 48
	// c size: 40
}
