/*
指针类型转化（不安全）

1. 任何类型的指针都可以被转化为 unsafe.Pointer；
2. unsafe.Pointer 可以被转化为任何类型的指针；
3. uintptr 可以被转化为 unsafe.Pointer；
4. unsafe.Pointer 可以被转化为 uintptr。
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	i := 10
	fmt.Println("i =", i) // i = 10
	var p *int = &i
	fmt.Printf("p = %p, type of %#v \n", p, p) // p = 0xc00001a098, type of (*int)(0xc00001a098)

	var fp *float32 = (*float32)(unsafe.Pointer(p))
	fmt.Printf("fp = %p, type of %#v\n", fp, fp) // fp = 0xc00001a098, type of (*float32)(0xc00001a098)
	*fp = *fp * 10
	fmt.Println("i =", i) // i = 100

}
