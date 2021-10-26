/*
内存缓存技术
*/
package main

import (
	"fmt"
	"time"
)

type FibonacciFunc1 func(int) int

const MAX = 50

// 通过预定义数组 fibs 保存已经计算过的斐波那契序号对应的数值
var fibs [MAX]int

// 通过递归函数实现斐波那契数列
func fibonacci1(n int) int {
	// 终止条件
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}

	// 递归公式
	num := fibonacci1(n-1) + fibonacci1(n-2)
	fibs[index] = num

	return num
}

// 斐波那契函数执行耗时计算
func fibonacciExecTime1(f FibonacciFunc1) FibonacciFunc1 {
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
	f := fibonacciExecTime1(fibonacci1)
	r1 := f(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

	n2 := 50
	r2 := f(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)
}

/*
1s = 10亿ns
1µs=1000ns

可以看到两者性能差不多在一个数量级上

======= 执行耗时：333ns ======
The 5th number of fibonacci sequence is 3
======= 执行耗时：542ns ======
The 50th number of fibonacci sequence is 7778742049
*/
