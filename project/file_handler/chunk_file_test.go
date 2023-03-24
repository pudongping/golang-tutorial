package file_handler

import (
	"log"
	"os"
	"testing"
)

func TestGenerateChunkFile(t *testing.T) {
	path, _ := os.Getwd()
	filePath := path + "/" + "haha.mp4"

	// 1MB
	var chunkSize int64 = 1 * 1024 * 1024

	GenerateChunkFile(filePath, chunkSize)
}

func TestMergeChunkFile(t *testing.T) {
	path, _ := os.Getwd()
	filePath := path + "/" + "haha.mp4"

	MergeChunkFile(filePath)
}

func TestCheckUniformity(t *testing.T) {

	newFile := "./merge_haha.mp4"
	oldFile := "./haha.mp4"

	is := CheckUniformity(newFile, oldFile)
	log.Printf("分片前后的文件是否具有一致性？ %v \n", is)

}
