package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("代码清理逻辑")
		if err := recover(); err != nil {
			fmt.Printf("recover 的信息为：%v \n", err)
		}
	}()

	var i = 1
	var j = 0
	if j == 0 {
		panic("除数不能为 0！")
	}

	// 剩下的代码不会被执行，因为 panic 注定会发生
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)

}

/*

代码清理逻辑
recover 的信息为：除数不能为 0！


*/
