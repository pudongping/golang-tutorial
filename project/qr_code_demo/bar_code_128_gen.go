package main

import (
	"image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func main() {
	// 要编码的文本
	text := "Hello World!"

	// 创建 Code128 条形码编码器
	enc := oned.NewCode128Writer()

	// 编码文本为 Code128 条形码图像
	img, _ := enc.Encode(text, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)

	// 创建文件保存条形码图像
	file, _ := os.Create("barcode.png")
	defer file.Close()

	// 将条形码图像编码为 PNG 格式并保存到文件
	_ = png.Encode(file, img)
}
