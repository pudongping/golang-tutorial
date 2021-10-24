package main

import "fmt"

func main() {

	// & => 取地址符，以获取该变量对应的内存地址
	// * => 间接引用符，以获取指针指向内存空间存储的变量值

	a := 100
	var ptr *int                   // 声明指针类型，表示指向存储 int 类型值的指针
	fmt.Println("ptr =", ptr)      // ptr = <nil>
	ptr = &a                       // ptr 本身是一个内存地址值，所以需要通过内存地址进行赋值（通过 &a 可以获取变量 a 所在的内存地址）
	fmt.Println("ptr =", ptr)      // ptr = 0xc00001a098  （这里每次打印的 ptr 值都有可能会不一样，因为存储变量 a 的内存地址在变动）
	fmt.Printf("ptr => %p\n", ptr) // ptr => 0xc0000ae008
	// 可以通过 *ptr 获取指针指向内存地址存储的变量值
	fmt.Println("*ptr =", *ptr) // *ptr = 100

	fmt.Println("===================")
	b := new(int)  // 通过内置函数 new 声明指针
	fmt.Println("b =", b)  // b = 0xc000118020
	*b = 200
	fmt.Println("*b =", *b)  // *b = 200

}
