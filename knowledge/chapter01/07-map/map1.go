package main

import "fmt"

func printMap(myCityMap map[string]string) {
	// 这里的参数其实是一个引用传递，也就意味着传的是指针

	// 遍历
	for key, value := range myCityMap {
		fmt.Println("key =", key, "value =", value)
	}

}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {

	// 声明一个 map
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 查找元素
	value, ok := cityMap["China"]
	if ok {
		fmt.Println("找到了 value =>", value) // 找到了 value => Beijing
	}

	fmt.Println("===========================")
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key =", key, "value =", value)
	}

	fmt.Println("===========================")

	// 只获取字典的键名
	for key := range cityMap {
		fmt.Println("key =", key)
	}

	// 删除 map 中的一个元素
	delete(cityMap, "China")

	// 修改 map 中的一个元素
	cityMap["USA"] = "DC"
	changeValue(cityMap)

	fmt.Printf("============= \n")

	printMap(cityMap)

}
