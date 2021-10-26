/*
通过尾递归优化递归函数性能
*/
package main

import (
	"fmt"
	"time"
)

type FibonacciFunc3 func(int) int

// 尾递归版本的斐波那契函数
func fibonacci3(n int) int {
	return fibonacciTail(n, 0, 1)
}

func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second)
}

// 斐波那契函数执行耗时计算
func fibonacciExecTime3(f FibonacciFunc3) FibonacciFunc3 {
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
	f := fibonacciExecTime3(fibonacci3)
	r1 := f(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

	n2 := 50
	r2 := f(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)
}

/*
1s = 10亿ns

======= 执行耗时：417ns ======
The 5th number of fibonacci sequence is 3
======= 执行耗时：417ns ======
The 50th number of fibonacci sequence is 7778742049
*/
