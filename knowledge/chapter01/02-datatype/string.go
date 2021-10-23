package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 转义字符
	results := "Search results for \"Golang\":\n" +
		"- Go\n" +
		"- Golang\n" +
		"- Golang Programming\n"
	/*
		The length of "Search results for "Golang":
		- Go
		- Golang
		- Golang Programming
		" is 64
	*/
	fmt.Printf("The length of \"%s\" is %d \n", results, len(results))

	// 推荐使用这种方式构建多行字符串，更加优雅
	results1 := `Search results for "Golang":
- Go
- Golang
- Golang Programming
`
	/*
		The length of "Search results for "Golang":
		- Go
		- Golang
		- Golang Programming
		" is 64
	*/
	fmt.Printf("The length of \"%s\" is %d \n", results1, len(results1))

	// 取字符串的第一个字符
	ch := results1[0]
	/*
		The first character of "Search results for "Golang":
		- Go
		- Golang
		- Golang Programming
		" is S.
	*/
	fmt.Printf("The first character of \"%s\" is %c. \n", results1, ch)

	// 切片区间是一个 **左闭右开** 的区间
	str := "hello, world"
	str1 := str[:5]              // 获取索引 5 （不含）之前的子串
	str2 := str[7:]              // 获取索引 7 （含）之后的子串
	str3 := str[0:5]             // 获取从索引 0 （含）到索引 5 （不含）之间的子串
	fmt.Println("str1 = ", str1) // str1 =  hello
	fmt.Println("str2 = ", str2) // str2 =  world
	fmt.Println("str3 = ", str3) // str3 =  hello

	str4 := str[:]
	fmt.Println("str4 = ", str4) // str4 =  hello, world

	fmt.Println("============ 优美的分割线 ==============")

	// 字符编码
	// 以字节数组的方式遍历
	msg := "Hello, 世界"
	n := len(msg) // 获取到的是 msg 的字节长度
	for i := 0; i < n; i++ {
		ch := msg[i] // 依据下标取字符串中的字符，ch 类型为 byte
		fmt.Println(i, ch, reflect.TypeOf(ch))
	}

	/*
		0 72 uint8
		1 101 uint8
		2 108 uint8
		3 108 uint8
		4 111 uint8
		5 44 uint8
		6 32 uint8
		7 228 uint8
		8 184 uint8
		9 150 uint8
		10 231 uint8
		11 149 uint8
		12 140 uint8
	*/

	fmt.Println("============ 优美的分割线 ==============")

	// 以 Unicode 字符遍历
	for i, ch := range msg {
		fmt.Println(i, string(ch), reflect.TypeOf(ch)) // 通过 range 遍历，ch 类型是 rune
	}

	/*
		0 H int32
		1 e int32
		2 l int32
		3 l int32
		4 o int32
		5 , int32
		6   int32
		7 世 int32
		10 界 int32
	*/

	fmt.Println("============ 优美的分割线 ==============")

}
