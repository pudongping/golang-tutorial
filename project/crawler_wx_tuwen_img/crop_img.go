package crawler_wx_tuwen_img

import (
	"bytes" // 用于处理字节缓冲区
	"fmt"   // 用于格式化输出
	"io"
	"net/http"      // 用于处理HTTP请求，这里主要用于MIME类型检测
	"os"            // 用于文件系统操作
	"path/filepath" // 用于文件路径操作
	"strings"
	"sync" // 用于并发控制

	"github.com/disintegration/imaging" // 第三方图像处理库，提供更高级的图像处理功能
)

// 定义一个全局的bufferPool，使用sync.Pool来重用字节缓冲区，减少内存分配和垃圾回收开销
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// isImage 函数检查给定文件是否为图片，通过读取文件的前512字节并检测其MIME类型实现
func isImage(file *os.File) bool {
	// 从bufferPool中获取一个字节缓冲区，并在函数返回时将其放回pool中
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	buffer.Reset()

	// 从文件中读取前512字节到缓冲区中
	_, err := buffer.ReadFrom(io.LimitReader(file, 512))
	if err != nil {
		return false
	}

	// 使用http.DetectContentType来检测缓冲区中数据的MIME类型
	contentType := http.DetectContentType(buffer.Bytes())
	// 如果MIME类型以"image/"开头，则认为文件是图片
	return strings.HasPrefix(contentType, "image/")
}

// processFile 函数处理单个图片文件，包括打开文件、检测是否为图片、读取图片、裁剪图片和保存图片
func processFile(path string, bottomPixel int, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()                // 在函数结束时通知WaitGroup，表示一个goroutine完成了工作
	defer func() { <-semaphore }() // 释放信号量，允许其他goroutine开始执行

	// 打开图片文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // 确保文件在函数返回时关闭

	// 检查文件是否为图片
	if !isImage(file) {
		fmt.Println(fmt.Sprintf("%+v 该文件不是图片", file.Name()))
		return
	}

	// 使用第三方库imaging打开并解码图片
	img, err := imaging.Open(path)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 使用imaging库裁剪图片，裁掉底部的80像素
	croppedImg := imaging.CropAnchor(img, img.Bounds().Dx(), img.Bounds().Dy()-bottomPixel, imaging.Top)

	// 将裁剪后的图片保存回原文件
	if err := imaging.Save(croppedImg, path); err != nil {
		fmt.Println("Error saving image:", err)
	} else {
		fmt.Println(fmt.Sprintf("裁剪并保存成功！ ==> %+v", path))
	}
}

// processImages 函数遍历指定目录及其子目录，寻找图片文件并并发处理它们
func processImages(rootDir string, bottomPixel int, concurrency int) {
	// 创建一个信号量通道，用于限制并发数量
	semaphore := make(chan struct{}, concurrency)

	var wg sync.WaitGroup // WaitGroup用于等待所有goroutine完成

	// 遍历目录中的所有文件和子目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件，则启动一个goroutine来处理它
		if !info.IsDir() {
			wg.Add(1)                                         // 增加WaitGroup的计数
			semaphore <- struct{}{}                           // 获取信号量，如果信号量用尽则阻塞，直到其他goroutine释放信号量
			go processFile(path, bottomPixel, &wg, semaphore) // 启动goroutine处理文件
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}

	wg.Wait() // 等待所有goroutine完成
}

func RunCropImg() {
	rootDir := "./download"                          // 设置要处理的目录路径
	concurrency := 10                                // 设置并发处理的goroutines数量
	bottomPixel := 65                                // 裁剪底部的65像素
	processImages(rootDir, bottomPixel, concurrency) // 开始处理图片
}
