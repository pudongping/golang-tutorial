/*
字典排序
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	testMap := make(map[string]int)
	testMap["first"] = 1
	testMap["second"] = 2
	testMap["third"] = 3

	for key, value := range testMap {
		fmt.Println("key =", key, "value =", value)
	}

	fmt.Println("============================")
	// 取所有的 key
	keys := make([]string, 0)
	for k := range testMap {
		keys = append(keys, k)
	}
	fmt.Println("keys =", keys)

	sort.Strings(keys) // 对键进行排序
	fmt.Println("Sorted map by key: ")
	for _, k := range keys {
		fmt.Println("k =", k, ", value =", testMap[k])
	}

	fmt.Println("============================")
	// 键值互换
	invMap := make(map[int]string)
	for key, value := range testMap {
		invMap[value] = key
	}
	// 取所有的 value
	values := make([]int, 0)
	for _, v := range testMap {
		values = append(values, v)
	}
	fmt.Println("values =", values)

	sort.Ints(values) // 对值进行排序
	fmt.Println("Sorted map by values: ")
	for _, v := range values {
		fmt.Println("k =", v, ", value =", invMap[v])
	}

}

/*
key = second value = 2
key = third value = 3
key = first value = 1
============================
keys = [first second third]
Sorted map by key:
k = first , value = 1
k = second , value = 2
k = third , value = 3
============================
values = [2 3 1]
Sorted map by values:
k = 1 , value = first
k = 2 , value = second
k = 3 , value = third
*/
