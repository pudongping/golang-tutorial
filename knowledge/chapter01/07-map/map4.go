/**
并发读写 map 案例
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Map 其是专门为 append-only 场景设计的，也就是适合读多写少的场景。这是他的优点之一。
若出现写多/并发多的场景，会导致 read map 缓存失效，需要加锁，冲突变多，性能急剧下降。这是他的重大缺点。
*/
var m sync.Map

func main() {
	// 写入
	data := []int{11, 22, 33, 44, 55}

	for i := 0; i <= 4; i++ {
		go func(i int) {
			// Store 存储并设置一个键的值
			m.Store(i, data[i])
		}(i)
	}

	time.Sleep(time.Second)

	// 读取
	// Load 返回存储在 map 中的键的值，如果没有值，则返回 nil。ok 结果表示是否在 map 中找到了值
	v, ok := m.Load(0)
	fmt.Printf("Load: %v, %v \n", v, ok)

	// 删除
	// Delete 删除某一个键的值
	m.Delete(1)

	// 读或写
	// LoadOrStore 如果存在的话，则返回键的现有值。否则，它存储并返回给定的值。如果值被加载，加载的结果为 true，如果被存储，则为 false
	v, ok = m.LoadOrStore(1, 123)
	fmt.Printf("LoadOrStore: %v, %v \n", v, ok)

	// 遍历
	// Range 递归调用，对 map 中存在的每个键和值依次调用闭包函数 f。如果 f 返回 false 就停止迭代。
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Range: %v, %v \n", key, value)

		return true
	})

}

/**
Load: 11, true
LoadOrStore: 123, false
Range: 0, 11
Range: 1, 123
Range: 3, 44
Range: 4, 55
Range: 2, 33
*/
