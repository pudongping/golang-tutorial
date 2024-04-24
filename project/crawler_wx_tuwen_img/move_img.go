package crawler_wx_tuwen_img

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var (
	reImg = regexp.MustCompile(`(?i)\.(jpg|jpeg|png|gif)$`)
)

type MoveImg struct {
	DirPath     string
	MaxNumImage uint
}

func (m *MoveImg) isImageFile(filename string) bool {
	// 支持 JPEG, PNG, GIF 等格式
	return reImg.MatchString(filename)
}

func genRandomNumber() int {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子确保每次执行的随机化不同
	// 生成1到100之间的随机数
	randomNumber := rand.Intn(100) + 1 // Intn(100)生成0到99之间的数，+1后变为1到100

	return randomNumber
}

func (m *MoveImg) randomizeImgFiles(files []string) map[string]string {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子确保每次执行的随机化不同
	rand.Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	// 创建一个临时映射，以避免直接覆盖文件
	tempMap := make(map[string]string)
	for index, file := range files {
		ext := filepath.Ext(file)
		// 重新生成新的文件名，避免修改文件名称时，出现命名冲突
		randomNumber := genRandomNumber()
		newName := fmt.Sprintf("%s/%d_%d%s", filepath.Dir(file), index+1, randomNumber, ext)
		tempMap[file] = newName
	}

	return tempMap
}

func (m *MoveImg) WalkAllImages() (imgChanges []map[string]string, err error) {
	// 使用filepath.Walk递归遍历目录
	err = filepath.Walk(m.DirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "Walk 出错：%s", path)
		}
		if info.IsDir() {
			// 因为这里只考虑了图片会放到文件夹下，所有只考虑文件夹下中的图片
			files, err := ioutil.ReadDir(path)
			if err != nil {
				return errors.Wrap(err, "读取目录出错")
			}
			var imageFiles []string
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if m.isImageFile(file.Name()) {
					// 收集符合条件的图片文件
					imageFiles = append(imageFiles, filepath.Join(path, file.Name()))
				}
			}
			// 打乱图片文件，并修改文件名称
			if len(imageFiles) > 1 {
				ret := m.randomizeImgFiles(imageFiles)
				imgChanges = append(imgChanges, ret)
			}
		}

		return nil
	})

	return
}

func (m *MoveImg) RenameImages(imgChanges []map[string]string) (errorSlices []error) {
	// 执行重命名
	for _, tempMap := range imgChanges {
		for original, newName := range tempMap {
			if err := os.Rename(original, newName); err != nil {
				errorSlices = append(errorSlices, errors.Wrapf(err, "%s 改名时，改成 %s 出错", original, newName))
			} else {
				log.Printf("图片改名 ==> %s ==> %s \n", original, newName)
			}
		}
	}

	return
}

func (m *MoveImg) DisassembleDir(imgChanges []map[string]string) error {
	for _, item := range imgChanges {
		if err := m.disassembleDirItem(item); err != nil {
			return errors.Wrap(err, "拆分目录时出现异常")
		}
	}

	return nil
}

func (m *MoveImg) getRandomImgPath(imgs map[string]string) string {
	values := make([]string, 0, len(imgs))
	for _, v := range imgs {
		values = append(values, v)
	}

	rand.Seed(time.Now().UnixNano())      // 初始化随机数生成器
	randomIndex := rand.Intn(len(values)) // 获取一个随机索引
	return values[randomIndex]
}

func (m *MoveImg) disassembleDirItem(imgs map[string]string) error {
	imageCount := len(imgs)
	if m.MaxNumImage <= 0 {
		return nil
	}
	if imageCount <= int(m.MaxNumImage) {
		return nil
	}

	var err error
	halfCount := imageCount / 2
	imageFiles := make([]string, 0, imageCount)
	for _, v := range imgs {
		imageFiles = append(imageFiles, v)
	}

	// 创建新的子目录
	parentPath := filepath.Dir(imageFiles[0]) // 随便取一个图片，从而获得父目录
	subDir1Name := filepath.Join(parentPath, strconv.Itoa(imageCount)+"_1")
	subDir2Name := filepath.Join(parentPath, strconv.Itoa(imageCount)+"_2")

	err = new(Tool).MkdirIfNotExist(subDir1Name)
	if err != nil {
		return errors.Wrap(err, "创建目录1出错")
	}
	err = new(Tool).MkdirIfNotExist(subDir2Name)
	if err != nil {
		return errors.Wrap(err, "创建目录2出错")
	}

	// 对图片进行排序并移动
	sort.Strings(imageFiles)
	for i, imageFile := range imageFiles {
		imgFileName := filepath.Base(imageFile)
		var dstPath string
		if i < halfCount {
			dstPath = filepath.Join(subDir1Name, imgFileName)
		} else {
			dstPath = filepath.Join(subDir2Name, imgFileName)
		}

		// 移动图片
		if err := os.Rename(imageFile, dstPath); err != nil {
			return errors.Wrapf(err, "%s ==> %s 移动失败 %v", imageFile, dstPath, err)
		} else {
			log.Printf("移动图片 ==> %s ==> %s \n", imageFile, dstPath)
		}

	}

	return nil
}

func RunMoveImg() {
	start := time.Now()

	dir := "./download"

	m := MoveImg{
		DirPath:     dir,
		MaxNumImage: 5, // 每个文件夹中最多可放的图片数量
	}
	log.Printf("开始找 %s 目录下所有的图片 \n", dir)
	imgChanges, err := m.WalkAllImages()
	if err != nil {
		fmt.Printf("Error: %+v \n", err)
		return
	}

	log.Println("开始打乱图片顺序并更改名称")
	errorSlices := m.RenameImages(imgChanges)
	for _, errorItem := range errorSlices {
		fmt.Printf("Error: %+v \n -------- \n\n", errorItem)
	}

	log.Println("拆分图片")
	if err := m.DisassembleDir(imgChanges); err != nil {
		fmt.Printf("Error: %+v \n", err)
	}

	castTime := time.Since(start)
	log.Printf("总耗时：%s", castTime.String())
}
