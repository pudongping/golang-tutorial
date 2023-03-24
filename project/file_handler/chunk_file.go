package file_handler

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GenerateChunkFile 生成分片文件
func GenerateChunkFile(filePath string, chunkSize int64) {
	relPath := filepath.Dir(filePath) + "/"

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("查看文件信息错误：%v \n", err)
	}

	// 计算分片个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("打开文件失败：%v \n", err)
	}

	block := make([]byte, chunkSize) // 每个分块
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		_, err := myFile.Seek(int64(i*int(chunkSize)), 0)
		if err != nil {
			log.Fatalf("出现错误：%v \n", err)
		}
		lastSize := fileInfo.Size() - int64(i*int(chunkSize)) // 还剩下的没有被分割的大小
		if lastSize < chunkSize {
			// 还没有被分片的尺寸小于每块需要被拆分的尺寸时
			block = make([]byte, lastSize) // 即最后一个分片的尺寸是不足 chunkSize 的
		}
		_, err = myFile.Read(block)
		if err != nil {
			log.Fatalf("读取失败：%v \n", err)
		}

		chunkFileName := relPath + fileInfo.Name() + "_" + strconv.Itoa(i) + ".chunk"
		f, err := os.OpenFile(chunkFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("新文件打开失败：%v \n", err)
		}
		_, err = f.Write(block)
		if err != nil {
			log.Fatalf("文件写入失败： %v \n", err)
		}
		if err = f.Close(); err != nil {
			log.Fatalf("文件关闭失败：%v \n", err)
		}

	}

	if err = myFile.Close(); err != nil {
		log.Fatalf("原始文件关闭失败：%v \n", err)
	}

}

// MergeChunkFile 合并分片文件
func MergeChunkFile(filePath string) {
	relPath := filepath.Dir(filePath) + "/"
	oldFileName := filepath.Base(filePath) // 旧的文件名称
	newFileName := relPath + "merge_" + oldFileName

	// 找出所有的分片（默认为，分片和旧文件在同一级目录下）
	var chunkNum int
	files, _ := ioutil.ReadDir(relPath)
	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), oldFileName+"_") {
			continue
		}
		log.Println("当前旧文件的分片文件为：", f.Name())
		chunkNum++
	}
	if chunkNum == 0 {
		log.Fatal("没有找到分片文件，不做合并操作")
	}

	myFile, err := os.OpenFile(newFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("打开新文件失败：%v \n", err)
	}

	// 开始合并文件
	for i := 0; i < chunkNum; i++ {
		chunkFileName := filePath + "_" + strconv.Itoa(i) + ".chunk"

		f, err := os.OpenFile(chunkFileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("打开分片文件失败：%v \n", err)
		}
		block, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("读取分片文件失败：%v \n", err)
		}

		_, err = myFile.Write(block)
		if err != nil {
			log.Fatalf("写入文件失败：%v \n", err)
		}
		if err := f.Close(); err != nil {
			log.Fatalf("关闭分片文件失败：%v \n", err)
		}

		log.Printf("开始删除分片文件： %s \n", chunkFileName)
		os.Remove(chunkFileName)

	}

	if err = myFile.Close(); err != nil {
		log.Fatalf("原始文件关闭失败：%v \n", err)
	}
}

// CheckUniformity 检验文件一致性
func CheckUniformity(newFile, oldFile string) bool {
	var rfBytes = func(fileName string) []byte {
		fl, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
		if err != nil {
			log.Fatalf("打开文件失败：%v \n", err)
		}
		b, err := ioutil.ReadAll(fl)
		if err != nil {
			log.Fatalf("读取文件失败： %v \n", err)
		}
		return b
	}

	file1 := rfBytes(newFile)
	file2 := rfBytes(oldFile)

	s1 := fmt.Sprintf("%x", md5.Sum(file1))
	s2 := fmt.Sprintf("%x", md5.Sum(file2))

	log.Printf("新文件的 md5 ==> %s \n", s1)
	log.Printf("旧文件的 md5 ==> %s \n", s2)

	return s1 == s2
}
