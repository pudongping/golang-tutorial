# init 函数与 import 导包

![import 导包执行流程](https://upload-images.jianshu.io/upload_images/14623749-006eb3b5910ec8e3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## import 导包

1. import _ "fmt"

给 fmt 包起一个别名，匿名，无法使用当前包的方法，但是会执行当前的包内部的 init() 方法

2. import aa "fmt"

给 fmt 包起一个别名 aa，可以使用 aa.Println() 来直接调用

3. import . "fmt"

将当前 fmt 包中的全部方法导入到当前本包的作用中，fmt 包中的全部方法可以直接使用 API 来调用，不需要 fmt.API 的方式来调用

## 导包的两种方式

```go

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

```