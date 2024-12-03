package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/corona10/goimagehash"
)

// compareImageSimilarity 比较图片相似度
func compareImageSimilarity(image1Path, image2Path string) error {
	file1, err := os.Open(image1Path)
	if err != nil {
		return fmt.Errorf("打开图片1失败：%v", err)
	}
	defer file1.Close()

	file2, err := os.Open(image2Path)
	if err != nil {
		return fmt.Errorf("打开图片2失败：%v", err)
	}
	defer file2.Close()

	// 加载图片
	img1, err := png.Decode(file1)
	if err != nil {
		return fmt.Errorf("加载图片1失败：%v", err)
	}

	img2, err := png.Decode(file2)
	if err != nil {
		return fmt.Errorf("加载图片2失败：%v", err)
	}

	// 生成平均哈希
	avgHash1, err := goimagehash.AverageHash(img1)
	if err != nil {
		return fmt.Errorf("生成图片1哈希失败：%v", err)
	}

	avgHash2, err := goimagehash.AverageHash(img2)
	if err != nil {
		return fmt.Errorf("生成图片2哈希失败：%v", err)
	}

	// 计算差异哈希
	diffHash1, err := goimagehash.DifferenceHash(img1)
	if err != nil {
		return fmt.Errorf("生成图片1差异哈希失败：%v", err)
	}

	diffHash2, err := goimagehash.DifferenceHash(img2)
	if err != nil {
		return fmt.Errorf("生成图片2差异哈希失败：%v", err)
	}

	// 计算汉明距离
	avgDistance, err := avgHash1.Distance(avgHash2)
	if err != nil {
		return fmt.Errorf("计算平均哈希距离失败：%v", err)
	}

	diffDistance, err := diffHash1.Distance(diffHash2)
	if err != nil {
		return fmt.Errorf("计算差异哈希距离失败：%v", err)
	}

	// 打印相似度
	fmt.Printf("平均哈希距离：%d\n", avgDistance)
	fmt.Printf("差异哈希距离：%d\n", diffDistance)

	// 判断相似程度
	if avgDistance == 0 && diffDistance == 0 {
		fmt.Println("两张图一样")
	} else if avgDistance <= 5 || diffDistance <= 5 {
		fmt.Println("图片高度相似")
	} else if avgDistance <= 10 || diffDistance <= 10 {
		fmt.Println("图片相似")
	} else {
		fmt.Println("图片差异较大")
	}

	return nil
}

func main() {
	err := compareImageSimilarity("img.png", "img_1.png")
	if err != nil {
		fmt.Println("图片对比出错：", err)
	}
}
