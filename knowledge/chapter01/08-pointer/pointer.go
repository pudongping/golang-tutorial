package main

import "fmt"

// 这种方式只是值做了交换，内存地址并没有做交换
/*func swap(a int, b int)  {
	var temp int
	temp = a
	a = b
	b = temp
}*/

func swap(pa *int, pb *int)  {
	//*pa = 10  // 表示将 *pa 内存对应的值修改成 10
	var temp int
	temp = *pa  // temp 其实直接等于 main 函数中的 a
	*pa = *pb  // 让 main 函数中的变量 b 直接赋值给 main 函数中的变量 a
	*pb = temp  // 然后再通过临时变量的方式将 main 函数中的变量 a 的内存地址指向 main 函数中的变量 b
}

func main()  {
	var a int = 10
	var b int = 20

	//swap(a, b)

	swap(&a, &b)  // &a 表示将 a 的内存地址传入进去

	fmt.Println("a =", a, "b =", b)  // a = 20 b = 10

	var p *int  // 声明一个指针类型的变量（指向 int 类型的变量）
	p = &a  // 将变量 a 的内存地址传给变量 p，如下所示，可见变量 p 已经指向了变量 a 的内存地址
	fmt.Println(&a)  // 0xc0000ae008
	fmt.Println(p)  // 0xc0000ae008

	var pp **int  // 二级指针（往上跳两次，先往上找到一级指针的内存地址，然后再往上找到变量的内存地址）
	pp = &p
	fmt.Println(&p)  // 0xc00018e020
	fmt.Println(pp)  // 0xc00018e020

}
