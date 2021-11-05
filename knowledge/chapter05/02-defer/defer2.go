package main

import "fmt"

func printError() {
	fmt.Println("兜底执行")
}

func main() {

	// 由于 defer 语句的执行时机和调用顺序，所以我们要尽量在函数/方法的前面定义它们，
	// 以免在后面编写代码时漏掉，尤其是运行时抛出错误，会中断后面代码的执行，也就感知不到后面的 defer 语句
	defer printError()
	defer func() {
		fmt.Println("除数不能为 0！")
	}()

	i := 1
	j := 1
	k := i / j

	fmt.Printf("%d / %d = %d\n", i, j, k)

}

/*
1 / 1 = 1
除数不能为 0！
兜底执行
*/
