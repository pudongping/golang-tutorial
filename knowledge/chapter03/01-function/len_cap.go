package main

import "fmt"

func main() {

	str := "golang"
	println(len(str)) // 6

	arr := [3]int{1, 2, 3}
	println("arr len is", len(arr), "cap is", cap(arr)) // arr len is 3 cap is 3

	slice := arr[1:]
	// slice len is 2 cap is 2 slice = [2/2]0xc000198008
	println("slice len is", len(slice), "cap is", cap(slice), "slice =", slice)
	fmt.Println("slice =", slice) // slice = [2 3]

	dict := map[string]int{"0": 1, "1": 2, "2": 3}
	// dict len is 3 dict = 0xc000180030
	println("dict len is", len(dict), "dict =", dict)
	// dict = map[0:1 1:2 2:3]
	fmt.Println("dict =", dict)

}
