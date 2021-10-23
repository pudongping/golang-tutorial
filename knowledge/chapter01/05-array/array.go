package main

import "fmt"

func main() {

	var a [8]byte           // 长度为8的数组，每个元素为一个字节
	var b [3][3]int         // 二维数组（9宫格）
	var c [3][3][3]float64  // 三维数组（立体的9宫格）
	var d = [3]int{1, 2, 3} // 声明时初始化
	var e = new([3]string)  // 通过 new 初始化

	fmt.Printf("a = %v, a = %#T\n", a, a) // a = [0 0 0 0 0 0 0 0], a = [8]uint8
	fmt.Printf("b = %v, b = %#T\n", b, b) // b = [[0 0 0] [0 0 0] [0 0 0]], b = [3][3]int
	fmt.Printf("c = %v, c = %#T\n", c, c) // c = [[[0 0 0] [0 0 0] [0 0 0]] [[0 0 0] [0 0 0] [0 0 0]] [[0 0 0] [0 0 0] [0 0 0]]], c = [3][3][3]float64
	fmt.Printf("d = %v, d = %#T\n", d, d) // d = [1 2 3], d = [3]int
	fmt.Printf("e = %v, e = %#T\n", e, e) // e = &[  ], e = *[3]string

	v1 := [5]int{2, 4, 6, 8}
	fmt.Printf("v1 = %v, v1 = %#T\n", v1, v1) // v1 = [2 4 6 8 0], v1 = [5]int
	// 通过语法糖省略数组长度来声明
	v2 := [...]string{"张三", "李四", "王五"}
	fmt.Printf("v2 = %v, v2 = %#T\n", v2, v2) // v2 = [张三 李四 王五], v2 = [3]string

}
