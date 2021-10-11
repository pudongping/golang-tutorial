/**
常量
 */
package main

import "fmt"

// iota 只能够配合 const 一起使用，iota 只有在 const 进行累加效果
const x = iota
const y = iota

// const 来定义枚举类型
const (
	// 可以在 const() 添加一个关键字 iota，每行的 iota 都会累加 1，第一行的 iota 的默认值是 0
	BEIJING = iota  // iota = 0
	SHANGHAI  // iota = 1
	SHENZHEN  // iota = 2
)

// 通过 iota 来定义公式
const (
	a, b = iota + 1, iota + 2  // iota = 0; a = iota + 1, b = iota + 2; a = 1, b = 2
	c, d					   // iota = 1; c = iota + 1, d = iota + 2; c = 2, d = 3
	e, f    				   // iota = 2; e = iota + 1, f = iota + 2; e = 3, f = 4
	g, h = iota * 2, iota * 3  // iota = 3; g = iota * 2, h = iota * 3; g = 6, h = 9
	i, k					   // iota = 4; i = iota * 2, k = iota * 3; i = 8, k = 12
)

func main()  {

	// 常量（只读属性）
	const length int = 10
	fmt.Println("length =", length)  // length = 10

	fmt.Println("BEIJING =", BEIJING)  // BEIJING = 0
	fmt.Println("SHANGHAI =", SHANGHAI)  // SHANGHAI = 1
	fmt.Println("SHENZHEN =", SHENZHEN)  // SHENZHEN = 2

	fmt.Println("a =", a, "b =", b)  // a = 1 b = 2
	fmt.Println("c =", c, "d =", d)  // c = 2 d = 3
	fmt.Println("e =", e, "f =", f)  // e = 3 f = 4
	fmt.Println("g =", g, "h =", h)  // g = 6 h = 9
	fmt.Println("i =", i, "k =", k)  // i = 8 k = 12

	fmt.Println("x =", x, "y =", y)  // x = 0 y = 0
}
