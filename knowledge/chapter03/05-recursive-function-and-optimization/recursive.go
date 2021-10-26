/*
递归实现斐波那契数列
*/
package main

import (
	"fmt"
	"time"
)

type FibonacciFunc func(int) int

// 通过递归函数实现斐波那契数列
func fibonacci(n int) int {
	// 终止条件
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// 递归公式
	return fibonacci(n-1) + fibonacci(n-2)
}

// 斐波那契函数执行耗时计算
func fibonacciExecTime(f FibonacciFunc) FibonacciFunc {
	return func(n int) int {
		start := time.Now()      // 起始时间
		num := f(n)              // 执行斐波那契函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("======= 执行耗时：%v ======\n", end)
		return num // 返回计算结果
	}
}

func main() {
	n1 := 5
	f := fibonacciExecTime(fibonacci)
	r1 := f(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

	n2 := 50
	r2 := f(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)
}

/*
1s = 10亿ns

======= 执行耗时：959ns ======
The 5th number of fibonacci sequence is 3
======= 执行耗时：37.833829084s ======
The 50th number of fibonacci sequence is 7778742049
*/
