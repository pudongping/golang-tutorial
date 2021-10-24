package main

import "fmt"

func main() {

	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				break
			}
			fmt.Println("num =", num)
		}
	}
	/*
		num = 1
		num = 2
		num = 4
		num = 5
		num = 7
		num = 8
	*/

	fmt.Println("================= 优美的分割线 =================")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				break
			} else {
				continue // 因为有了 continue，因此下面不会有任何输出
			}
			fmt.Println("num1 =", num)
		}
	}

	fmt.Println("================= 优美的分割线 =================")

ITERATOR1: // 标签语句通过 `标签加上冒号 :` 进行声明
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				break ITERATOR1 // 直接跳转到 ITERATOR1 标签对应的位置
			}
			fmt.Println("num2 =", num)
		}
	}
	/*
		num2 = 1
		num2 = 2
	*/

	fmt.Println("================= 优美的分割线 =================")

ITERATOR2: // 标签语句通过 `标签加上冒号 :` 进行声明
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				continue ITERATOR2 // 直接跳转到 ITERATOR1 标签对应的位置
			}
			fmt.Println("num3 =", num)
		}
	}
	/*
		num3 = 1
		num3 = 2
		num3 = 4
		num3 = 5
		num3 = 7
		num3 = 8
	*/

	fmt.Println("================= 优美的分割线 =================")

	// goto 语句
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				goto EXIT
			}
			fmt.Println("num4 =", num)
		}
	}

EXIT:
	fmt.Println("到了 EXIT 标签")
	/*
		num4 = 1
		num4 = 2
		到了 EXIT 标签
	*/

}
