/*
指针运算实现（不安全）
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr := [3]int{1, 2, 3}
	fmt.Println("arr =", arr) // arr = [1 2 3]
	ap := &arr
	fmt.Println("ap =", ap) // ap = &[1 2 3]

	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0]))) // 获取数组 arr 第二个元素的内存地址
	// 打印数组 arr 第二个元素的内存地址以及获取指针指向内存空间存储的变量值
	fmt.Println("sp =", sp, "&sp =", *sp) // sp = 0xc0000b6008 &sp = 2
	*sp += 44                             // 修改数组 arr 第二个元素的变量值
	fmt.Println("arr =", arr)             // arr = [1 46 3]

	/*
	   这样一来，就可以绕过 Go 指针的安全限制，实现对指针的动态偏移和计算了，
	   这会导致即使发生数组越界了，也不会报错，而是返回下一个内存地址存储的值，
	   这就破坏了内存安全限制，所以这也是不安全的操作，
	   我们在实际编码时要尽量避免使用，必须使用的话也要非常谨慎。
	*/

}
