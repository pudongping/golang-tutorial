/*
defer 的执行顺序
通过下面的例子可以得出结论：
defer 是类似于栈结构（先入后出，后入先出）的执行顺序，并且 defer 执行时机是在一个函数执行完毕之后才会去执行的
类似于 php 的 __destruct 析构函数，在函数执行结束后自动执行
 */
package main

import "fmt"

func func1()  {
	fmt.Println("func1")
}

func func2()  {
	fmt.Println("func2")
}

func func3()  {
	fmt.Println("func3")
}

func main()  {
	defer func1()
	defer func2()
	defer func3()

	fmt.Println("run main func")

}

/*
run main func
func3
func2
func1
*/