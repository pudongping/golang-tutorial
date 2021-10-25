/*
通过高阶函数实现装饰器模式
*/
package main

import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数 1 （算术运算）
func multiply1(a, b int) int {
	return a * b
}

// 乘法运算函数 2 （位运算）
func multiply2(a, b int) int {
	return a << b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()      // 起始时间； time.Now() 获取当前系统时间
		c := f(a, b)             // 执行乘法运算函数
		end := time.Since(start) // 函数执行完毕耗时； time.Since(start) 计算从 start 到现在经过的时间
		fmt.Printf("===== 执行耗时：%v ====== \n", end)
		return c // 返回计算结果
	}
}

func main() {
	a := 2
	b := 8

	fmt.Println("算术运算；")
	// 通过修饰器调用乘法函数，返回的是一个匿名函数
	docorator1 := execTime(multiply1)
	// 执行修饰器返回函数
	c := docorator1(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, c) // 2 * 8 = 16

	fmt.Println("位运算：")
	docorator2 := execTime(multiply2)
	a = 1
	b = 4
	c = docorator2(a, b)
	fmt.Printf("%d << %d = %d\n", a, b, c)

	/*
		算术运算；
		===== 执行耗时：292ns ======
		2 * 8 = 16
		位运算：
		===== 执行耗时：166ns ======
		1 << 4 = 16
	*/
}
