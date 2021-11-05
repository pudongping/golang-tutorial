/*
无论是 Go 语言底层抛出 panic，还是我们在代码中显式抛出 panic，处理机制都是一样的：
当遇到 panic 时，Go 语言会中断当前协程（即 main 函数）后续代码的执行，
然后执行在中断代码之前定义的 defer 语句（按照先入后出的顺序），最后程序退出并输出 panic 错误信息，以及出现错误的堆栈跟踪信息
*/
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("代码清理逻辑")
	}()

	var i = 1
	var j = 0
	if j == 0 {
		panic("除数不能为 0！")
	}

	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)

}

/*
代码清理逻辑
panic: 除数不能为 0！

goroutine 1 [running]:
main.main()
        /Users/pudongping/glory/codes/golang/golang-tutorial/knowledge/chapter05/03-panic-and-recover/panic.go:18 +0x5b
exit status 2

这里我们可以看到，是先走到了 defer，然后再走到 panic

*/
