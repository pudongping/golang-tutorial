package main

import (
	"fmt"
	"strconv"
)

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
	}

	// 用户年龄累加的结果为：60
	fmt.Printf("用户年龄累加的结果为：%d\n", ageSum(users))

}

func ageSum(users []map[string]string) int {
	var sum int
	for _, user := range users {
		num, _ := strconv.Atoi(user["age"])
		sum += num
	}
	return sum
}
