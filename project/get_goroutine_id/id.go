package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"

	"github.com/petermattis/goid"
)

// GetGoroutineID 返回当前 Goroutine 的 ID
// 通过 runtime.Stack 获取当前 Goroutine 的栈信息，然后提取出 Goroutine ID
// 这种方式可以获取到当前 Goroutine 的 ID，但是性能较差
func GetGoroutineID() uint64 {
	var buf [64]byte
	// runtime.Stack(buf[:], false) 会将当前 Goroutine 的栈信息写入 buf 中
	// 第二个参数是 false 表示只获取当前 Goroutine 的栈信息，如果为 true 则会获取所有 Goroutine 的栈信息
	n := runtime.Stack(buf[:], false)
	stack := string(buf[:n])
	// fmt.Println("========")
	// fmt.Println(stack)
	// fmt.Println()
	// stack 样例: "goroutine 7 [running]:\n..."
	// 提取 goroutine 后面的数字
	fields := bytes.Fields([]byte(stack))
	id, err := strconv.ParseUint(string(fields[1]), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("无法解析 Goroutine ID: %v", err))
	}
	return id
}

// goid 库使用了 C 和 汇编来获取 Goroutine ID，性能更好
func GetGoroutineID1() int64 {
	id := goid.Get()
	return id
}

func main() {
	fmt.Printf("Main Goroutine ID: %d\n", GetGoroutineID1())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Child [%d] Goroutine ID: [%d]\n", i, GetGoroutineID1())
		}()
	}

	wg.Wait()
}

// $ go run id.go
// 打印结果类似：
// Main Goroutine ID: 1
// Child [9] Goroutine ID: [15]
// Child [2] Goroutine ID: [8]
// Child [0] Goroutine ID: [6]
// Child [5] Goroutine ID: [11]
// Child [8] Goroutine ID: [14]
// Child [1] Goroutine ID: [7]
// Child [3] Goroutine ID: [9]
// Child [4] Goroutine ID: [10]
// Child [6] Goroutine ID: [12]
// Child [7] Goroutine ID: [13]
