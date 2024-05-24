package main

// 这里的包必须在 $GOPATH 下才行
import (
	//"./lib1"
	//"./lib2"
	//_ "./lib1"  // 当我们想导入 lib1 的包，但是却又不想使用 lib1 包里面的接口时，我们可以采用在导包前面加上 `_` 匿名导包的形式导入，此时会执行 lib1 的 init 方法
	//mylib2 "./lib2"  // 别名导包的方式
	//. "./lib2"  // 也可以通过在导包之前加上一个 `.` 号，将此包中的所有接口导入到当前包中，但是不推荐这么使用，因为可能会有同名方法会被覆盖
	"go-tutorial/knowledge/chapter04/09-init/lib1"
	"go-tutorial/knowledge/chapter04/09-init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
	//mylib2.Lib2Test()  // 别名导包时，使用接口时
}

/*lib1 ===> init() ...
lib2 ===> init() ...
Lib1Test() ...
Lib2Test() ...
*/
