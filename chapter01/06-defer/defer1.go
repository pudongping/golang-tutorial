/*
defer 和 return 谁先谁后
通过下面的例子可以得出结论：
return 语句要早于 defer 语句
 */
package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main()  {
	returnAndDefer()
}

/*
return func called ...
defer func called ...
*/