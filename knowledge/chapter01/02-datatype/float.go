package main

import "fmt"

func main() {

	floatValue1 := 0.01 // 对于浮点类型需要被自动推导的变量，其类型将被自动设置为 float64，而不管赋值给它的数字是否是用 32 位长度表示的
	floatValue2 := 0.2
	floatValue3 := floatValue1 + floatValue2
	floatValue4 := 1.1E-10
	fmt.Println("floatValue3 =", floatValue3) // floatValue3 = 0.21000000000000002
	fmt.Println("floatValue4 =", floatValue4) // floatValue4 = 1.1e-10

}
