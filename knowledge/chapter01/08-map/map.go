package main

import "fmt"

func main() {

	// ====> 第一种声明 map 的方式

	// 声明 myMap1 是一种 map 类型，key 是 string，value 也是 string
	var myMap1 map[string]string
	// myMap1 是一个空 map
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空 map")
	}

	// 在使用 map 前，需要先用 make 给 map 分配数据空间，这里分配 10 个空间
	myMap1 = make(map[string]string, 10)
	// 也可以在声明变量的时候直接分配空间
	// var myMap1 = make(map[string]string, 10)
	myMap1["name"] = "alex"
	myMap1["age"] = "18"
	myMap1["city"] = "Shanghai"
	// map[age:18 city:Shanghai name:alex]
	fmt.Println(myMap1)

	// ====> 第二种声明 map 的方式
	// 可以不给 map 分配空间，这样就采用动态空间
	myMap2 := make(map[int]string)
	myMap2[12] = "harry"
	myMap2[14] = "23"
	myMap2[18] = "Chongqing"
	// map[12:harry 14:23 18:Chongqing]
	fmt.Println(myMap2)

	// ====> 第三种声明 map 的方式
	myMap3 := map[string]string{
		"first":  "php",
		"second": "python",
		"third":  "golang",
	}
	// map[first:php second:python third:golang]
	fmt.Println(myMap3)

}
