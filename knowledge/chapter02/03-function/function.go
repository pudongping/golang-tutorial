/*
函数
*/
package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("---- foo1 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return c
}

// 返回多个返回值，匿名的返回值
func foo2(a string, b int) (int, int) {
	fmt.Println("---- foo2 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 666, 777
}

// 返回多个返回值，有行参名称的
func foo3(a string, b bool) (r1 int, r2 int)  {
	fmt.Println("---- foo3 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 在未赋值之前打印需要返回的结果参数
	// r1 r2 属于 foo3 的行参，初始化默认的值是 0
	// r1 r2 作用域空间是 foo3 整个函数体的 {} 空间
	fmt.Println("foo3 r1 =", r1)
	fmt.Println("foo3 r2 =", r2)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b bool) (r1, r2 int)  {
	fmt.Println("---- foo4 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 给有名称的返回值变量赋值
	r1 = 3000
	r2 = 4000

	return
}

func main() {
	c := foo1("alex", 555)
	fmt.Println("c =", c)

	ret1, ret2 := foo2("harry", 88)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo3("mark", false)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo4("jack", true)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}

/*
---- foo1 ----
a = alex
b = 555
c = 100
---- foo2 ----
a = harry
b = 88
ret1 = 666 ret2 = 777
---- foo3 ----
a = mark
b = false
foo3 r1 = 0
foo3 r2 = 0
ret1 = 1000 ret2 = 2000
---- foo4 ----
a = jack
b = true
ret1 = 3000 ret2 = 4000
*/