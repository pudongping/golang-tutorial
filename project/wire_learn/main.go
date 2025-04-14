package main

import (
	"context"
	"fmt"
)

func main() {
	cObj, err := InitializeC(context.Background(), "小明", 18)
	if err != nil {
		fmt.Println("初始化失败:", err)
		return
	}
	fmt.Println("初始化成功 ----> ", cObj.Content)
}
