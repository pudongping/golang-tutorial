package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("关注博主的人最帅！", qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		fmt.Println("生成二维码失败:", err)
	}
	fmt.Println("二维码已生成: qrcode.png")
}
