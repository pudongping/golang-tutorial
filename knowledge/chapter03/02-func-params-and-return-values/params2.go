/*
变长参数
*/
package main

import "fmt"

func myfunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println("number = ", number)
	}
}

func main() {

	myfunc(1, 2, 3, 4, 5)
	/*
		number =  1
		number =  2
		number =  3
		number =  4
		number =  5
	*/

	fmt.Println("=========== 优美的分割线 =============")
	// 还支持传递一个切片
	slice := []int{11, 22, 33, 44, 55}
	myfunc(slice...)
	/*
		number =  11
		number =  22
		number =  33
		number =  44
		number =  55
	*/

	fmt.Println("=========== 优美的分割线 =============")
	myfunc(slice[2:4]...)
	/*
		number =  33
		number =  44
	*/

}
