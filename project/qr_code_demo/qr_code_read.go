package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // 支持 JPEG 格式
	_ "image/png"  // 支持 PNG 格式
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	// 1. 打开二维码图片
	file, err := os.Open("qrcode.png") // 替换为你的二维码图片路径
	if err != nil {
		fmt.Println("无法打开图片:", err)
		return
	}
	defer file.Close()

	// 2. 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("无法解码图片:", err)
		return
	}

	// 3. 创建二维码解码器
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Println("无法创建二维码位图:", err)
		return
	}

	// 4. 识别二维码
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		fmt.Println("二维码识别失败:", err)
		return
	}

	// 5. 输出二维码内容
	fmt.Println("二维码内容:", result.GetText())
}
