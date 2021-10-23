package main

import "fmt"

func main() {

	// 通过二维数组生成九九乘法表
	var multi [9][9]string

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%d*%d=%d", n2, n1, n1*n2)
		}
	}

	fmt.Printf("%#v\n", multi)
	fmt.Println("=================================")

	// 打印
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) // 位宽为 8，左对齐
		}
		fmt.Println() // 打印换行
	}

}

/*
[9][9]string{[9]string{"1*1=1", "", "", "", "", "", "", "", ""}, [9]string{"1*2=2", "2*2=4", "", "", "", "", "", "", ""}, [9]string{"1*3=3", "2*3=6", "3*3=9", "", "", "", "", "", ""}, [9]string{"1*4=4", "2*4=8", "3*4=12", "4*4=16", "", "", "", "", ""}, [9]string{"1*5=5", "2*5=10", "3*5=15", "4*5=20", "5*5=25", "", "", "", ""}, [9]string{"1*6=6", "2*6=12", "3*6=18", "4*6=24", "5*6=30", "6*6=36", "", "", ""}, [9]string{"1*7=7", "2*7=14", "3*7=21", "4*7=28", "5*7=35", "6*7=42", "7*7=49", "", ""}, [9]string{"1*8=8", "2*8=16", "3*8=24", "4*8=32", "5*8=40", "6*8=48", "7*8=56", "8*8=64", ""}, [9]string{"1*9=9", "2*9=18", "3*9=27", "4*9=36", "5*9=45", "6*9=54", "7*9=63", "8*9=72", "9*9=81"}}
=================================
1*1=1
1*2=2   2*2=4
1*3=3   2*3=6   3*3=9
1*4=4   2*4=8   3*4=12  4*4=16
1*5=5   2*5=10  3*5=15  4*5=20  5*5=25
1*6=6   2*6=12  3*6=18  4*6=24  5*6=30  6*6=36
1*7=7   2*7=14  3*7=21  4*7=28  5*7=35  6*7=42  7*7=49
1*8=8   2*8=16  3*8=24  4*8=32  5*8=40  6*8=48  7*8=56  8*8=64
1*9=9   2*9=18  3*9=27  4*9=36  5*9=45  6*9=54  7*9=63  8*9=72  9*9=81
*/
