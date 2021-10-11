package main  // 程序的包名

/**
两种方式都可以导入
import "fmt"
import "time"
 */
import (
	"fmt"
	"time"
)

// main 函数
func main() {  // 函数的 `{` 一定是需要和函数名在同一行的，否则编译会错误
	// golang 中的表达式，加 `;` 和不加分号都可以，建议不加
	fmt.Println("hello Go!")
	time.Sleep(1 * time.Second)
}
