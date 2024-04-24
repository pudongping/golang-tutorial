package crawler_wx_tuwen_img

import (
	"fmt"
	"testing"
)

func TestRunSpiderImg(t *testing.T) {
	// 抓取图片
	err := RunSpiderImg()
	if err != nil {
		fmt.Println(err)
	}
}

func TestRunCropImg(t *testing.T) {
	// 裁剪图片
	RunCropImg()
}

func TestRunMoveImg(t *testing.T) {
	// 打乱并移动图片
	RunMoveImg()
}
