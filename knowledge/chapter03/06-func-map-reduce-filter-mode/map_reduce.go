package main

import (
	"fmt"
	"strconv"
)

// Map 函数
// 将字典类型切片转化为一个字符串类型切片
func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, item := range items {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

// Reduce 函数
// 再将转化后的切片元素转化为整型后进行累加
func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, item := range items {
		sum += f(item)
	}
	return sum
}

// Filter 函数
func itemsFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
	newSlice := make([]map[string]string, len(items))
	for _, item := range items {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func main() {

	// 定义一个字典类型用户切片
	// key 为 string，value 也为 string 的 map
	users := []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
		{
			"name": "小三",
			"age":  "-20",
		},
		{
			"name": "小王",
			"age":  "15",
		},
	}

	validUsers := itemsFilter(users, func(user map[string]string) bool {
		age, ok := user["age"]
		if !ok {
			return false
		}
		intAge, err := strconv.Atoi(age)
		if err != nil {
			return false
		}
		if intAge < 18 || intAge > 35 {
			return false
		}
		return true
	})

	ageSlice := mapToString(validUsers, func(user map[string]string) string {
		return user["age"]
	})

	fmt.Printf("ageSlice = %#v\n", ageSlice)

	sum := fieldSum(ageSlice, func(age string) int {
		intAge, _ := strconv.Atoi(age)
		return intAge
	})

	// 用户年龄累加的结果为：60
	fmt.Printf("用户年龄累加的结果为：%d\n", sum)
}
