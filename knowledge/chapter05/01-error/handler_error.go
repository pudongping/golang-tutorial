package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func add(a, b int) (c int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	a *= 2
	b *= 3
	c = a + b
	return
}

func main() {
	// 通过引入 os 包读取命令行参数
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s num1 num2\n", filepath.Base(os.Args[0]))
		return
	}
	x, _ := strconv.Atoi(os.Args[1]) // 转为整型
	y, _ := strconv.Atoi(os.Args[2])
	// 通过多返回值捕获函数调用过程中可能的错误信息
	z, err := add(x, y)
	// 通过 [卫述语句] 处理后续业务逻辑
	if err != nil {
		// 这里直接传入了 err 对象实例，因为 Go 底层会自动调用 err 实例上的 Error() 方法返回错误信息
		// 并将其打印出来，就像普通类的 String() 方法一样
		fmt.Println(err)
	} else {
		fmt.Printf("add(%d, %d) = %d\n", x, y, z)
	}
}

/*
go run handler_error.go -1 2
output: 只支持非负整数相加

go run handler_error.go 1 2
add(1, 2) = 8

*/
