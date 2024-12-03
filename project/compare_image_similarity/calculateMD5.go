package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// calculateMD5 计算文件的MD5值
// 为图片颁发独一无二的“身份证”
func calculateMD5(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件：%v", err)
	}
	defer file.Close()

	// 创建MD5哈希器
	hash := md5.New()

	// 将文件内容复制到哈希器中
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("计算MD5时出错：%v", err)
	}

	// 返回MD5字符串
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
	// 比较两张图片
	md51, err := calculateMD5("img.png")
	if err != nil {
		fmt.Println("图片1处理失败:", err)
		return
	}

	md52, err := calculateMD5("img_1.png")
	if err != nil {
		fmt.Println("图片2处理失败:", err)
		return
	}

	// 对比结果
	if md51 == md52 {
		fmt.Println("图片完全相同")
	} else {
		fmt.Println("图片不同")
	}
}
