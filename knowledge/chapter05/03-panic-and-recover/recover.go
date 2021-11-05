/*
我们可以通过 recover() 函数对 panic 进行捕获和处理，从而避免程序崩溃然后直接退出，而是继续可以执行后续代码。
由于执行到抛出 panic 的问题代码时，会中断后续其他代码的执行，
所以，显然这个 panic 的捕获应该放到 defer 语句中完成，才可以在抛出 panic 时通过 recover 函数将其捕获，
defer 语句执行完毕后，会退出抛出 panic 的当前函数，回调调用它的地方继续后续代码的执行。
可以类比为 panic、recover、defer 组合起来实现了传统面向对象编程异常处理的 try、catch、finally 功能。
*/
package main

import "fmt"

func devide() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()

	i, j := 1, 0
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)

}

func main() {
	devide()
	fmt.Println("devide 方法调用完毕，回到 main 函数")
}

/*
Runtime panic caught: runtime error: integer divide by zero
devide 方法调用完毕，回到 main 函数
*/
