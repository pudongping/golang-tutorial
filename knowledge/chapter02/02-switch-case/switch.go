package main

import "fmt"

func main() {

	// 不能把变量 score 放到 switch 关键字后面，否则会报错
	score := 100
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80 && score < 90:
		fmt.Println("Grade: B")
	case score >= 70 && score < 80:
		fmt.Println("Grade: C")
	case score >= 60 && score < 70:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// 只有在与 case 分支值判等的时候，才可以将变量放到 switch 关键字后面
	score1 := 70 // 这里会输出 Grade: C
	switch score1 {
	case 90, 100:
		fmt.Println("Grade: A")
	case 80:
		fmt.Println("Grade: B")
	case 70:
		fallthrough // 如果想要继续执行后续分支的代码，这里就相当于合并了 case 70 和 case 75 这两个分支语句，如果 score1 等于 70 的话，会直接打印出 `Grade: C`
	case 75:
		fmt.Println("Grade: C")
	case 60: // 这里当 score1 == 60 时，不会有任何输出，因为 Go 语言中会认为这里是一个空语句，会直接退出
	case 65:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

}
