package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// getRootPath 获取项目根目录
func getRootPath() string {

	// 第一种方式：获取当前执行程序所在的绝对路径
	// 这种仅在 `go build` 时，才可以获取正确的路径
	// 获取当前执行的二进制文件的全路径，包括二进制文件名
	// eg: exePath = "/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T/go-build265586886/b001/exe/main"
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// eg: rootPathByExecutable = "/private/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T/go-build265586886/b001/exe"
	rootPathByExecutable, _ := filepath.EvalSymlinks(filepath.Dir(exePath))

	// 第二种方式：获取当前执行文件绝对路径
	// 这种方式在 `go run` 和 `go build` 时，都可以获取到正确的路径
	// 但是交叉编译后，执行的结果是错误的结果
	var rootPathByCaller string
	// eg: filename = "/Users/pudongping/glory/codes/golang/gin-biz-web-api/main.go"
	_, filename, _, ok := runtime.Caller(0)
	// eg: rootPathByCaller = "/Users/pudongping/glory/codes/golang/gin-biz-web-api"
	if ok {
		rootPathByCaller = path.Dir(filename)
	}

	// 可以通过 `echo $TMPDIR` 查看当前系统临时目录
	// eg: tmpDir = "/private/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T"
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())

	// 对比通过 `os.Executable()` 获取到的路径是否与 `TMPDIR` 环境变量设置的路径相同
	// 相同，则说明是通过 `go run` 命令启动的
	// 不同，则是通过 `go build` 命令启动的
	if strings.Contains(rootPathByExecutable, tmpDir) {
		return rootPathByCaller
	}

	return rootPathByExecutable
}

func main() {
	// 当然还有其他的方式，比如可以自己先在系统里设置诸如 `PROJECT_ROOT_DIR` 之类的环境变量，将
	// 根目录放到环境变量中，然后再在程序中通过 `os.Getenv("PROJECT_ROOT_DIR")` 也行
	fmt.Println(getRootPath())
}
