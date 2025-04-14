package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// 模拟耗时操作，例如从数据库中获取数据
func fetchData(param string, id int) (string, error) {
	// 模拟操作耗时  1  秒
	time.Sleep(time.Second)
	// 当前时间（精确到纳秒）
	fmt.Printf("当前时间 %d \n", time.Now().UnixNano())
	return fmt.Sprintf("id %d 获取的数据： %s", id, param), nil
}

func main() {
	// 定义 singleflight.Group 对象
	var group singleflight.Group
	// 使用 WaitGroup 等待所有 goroutine 执行完毕
	var wg sync.WaitGroup

	// 模拟  5  个并发请求使用相同的 key "resource"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// group.Do 方法合并相同 key 的请求，只执行一次 fetchData
			value, err, shared := group.Do("resource", func() (interface{}, error) {
				// 这里传入实际需要执行的函数
				return fetchData("example", id)
			})
			if err != nil {
				fmt.Printf("请求 %d 出错： %v\n", id, err)
				return
			}
			// shared 表示结果是否为共享
			fmt.Printf("请求 %d 得到结果： [%v]  ，是否共享： %v\n", id, value, shared)
		}(i)
	}

	wg.Wait()
}
