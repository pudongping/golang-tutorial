/*
键值对调
*/
package main

import "fmt"

func main() {

	testMap := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	for key, value := range testMap {
		fmt.Println("key =", key, "value =", value)
	}

	fmt.Println("=========================")

	invMap := make(map[int]string)
	for key, value := range testMap {
		invMap[value] = key
	}
	for key, value := range invMap {
		fmt.Println("key =", key, "value =", value)
	}

}

/*
key = first value = 1
key = second value = 2
key = third value = 3
=========================
key = 3 value = third
key = 1 value = first
key = 2 value = second
*/
