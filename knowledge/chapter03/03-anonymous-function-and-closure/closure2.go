/*
将匿名函数作为函数返回值
*/
package main

import "fmt"

func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

func main() {

	// 此时返回的是匿名函数
	addFunc := deferAdd(1, 2)
	fmt.Println("========== 优美的分割线 ============")
	// 这里才会真正执行加法操作
	fmt.Println("addFunc =", addFunc())

	/*
	   ========== 优美的分割线 ============
	   addFunc = 3
	*/

}
