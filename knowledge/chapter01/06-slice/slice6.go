package main

import (
	"fmt"
	"math"
)

func main() {

	// var dataSlice []int = foo()
	// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	// for i, d := range dataSlice {
	//    interfaceSlice[i] = d
	// }

	slice1 := []string{"a1", "a2", "a3"}
	abc := CutSlice(slice1, 2)
	fmt.Printf("abc ===> %v type ==> %T \n", abc, abc)
	// abc ===> [[a1 a2] [a3 a4] [a5 a6] [a7]] type ==> [][]string

	// everyTime := 2
	// times := math.Ceil(float64(len(slice1)) / float64(everyTime))
	//
	// var slice2 [][]string
	// for i := 0; i < int(times); i ++ {
	// 	var slice3 []string
	// 	start := i * everyTime
	// 	end := start + everyTime
	//
	// 	fmt.Println("i ==> ", i)
	// 	if i == (int(times) - 1) {
	// 		fmt.Println("来来没 ==> ", i)
	// 		end = len(slice1)
	// 	}
	//
	// 	slice3 = slice1[start:end]
	// 	slice2 = append(slice2, slice3)
	// 	fmt.Printf("start ===> %d end ===> %d\n", start, end)
	//
	// }

}

// 将切片按照指定步长做切割
func CutSlice(s []string, l int) [][]string {

	if len(s) == 0 {
		return nil
	}

	times := math.Ceil(float64(len(s)) / float64(l))
	var s1 [][]string
	for i := 0; i < int(times); i++ {
		var s2 []string
		start := i * l
		end := start + l

		if i == (int(times) - 1) {
			end = len(s)
		}

		s2 = s[start:end]
		s1 = append(s1, s2)

	}

	return s1
}
