/*
使用多返回值的特性实现报错
*/
package main

import (
	"errors"
	"fmt"
)

func add1(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}

	*a *= 2
	*b *= 3
	return *a + *b, nil
}

// 命名返回值
func add2(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}

	*a *= 2
	*b *= 3
	c = *a + *b
	return
}

func main() {
	x, y := -1, 2
	z, err := add1(&x, &y)
	if err != nil {
		fmt.Println(err.Error()) // 只支持非负整数相加
		return
	}

	fmt.Printf("add1(%d, %d) = %d\n", x, y, z)

	fmt.Println("============== 优美的分割线 ==================")

	n, err := add2(&x, &y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("add2(%d, %d) = %d\n", x, y, n)

}
