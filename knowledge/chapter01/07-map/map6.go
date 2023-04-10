package main

// 在多个 goroutine 中共享 map 时，可以使用以下几种方式来保证 map 的安全：
//
// 1. 互斥锁：使用 sync.Mutex 可以对对 map 进行加锁和解锁操作，以避免多个 goroutine 同时对 map 进行读写操作，造成数据不一致的问题。
//
// 2. 读写锁：使用 sync.RWMutex 可以支持多个读操作并只限制一个写操作，以提高并发读取性能。
//
// 3. 使用 sync.Map 包

import (
	"fmt"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var mp map[string]int
var mm map[int64]int64
var lock sync.Mutex     // 互斥锁
var rwLock sync.RWMutex // 读写锁

// 互斥锁：
// 适用于写操作很少，读操作很多的场景。这种方法可以避免频繁锁/解锁带来的性能影响
func mapSyncMutex() {
	mp = make(map[string]int)

	var wg sync.WaitGroup

	// 如果多个 goroutine 同时对同一个 map 进行读写操作，它们可能会导致数据不一致，产生冲突
	// 启动多个 goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				key := "key"
				lock.Lock()
				mp[key]++
				lock.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	spew.Dump(mp)
}

// 读写锁：
// 适用于读写操作都比较多的场景。这种方法能够保证读操作的并发性，同时避免写操作带来的性能影响
func mapRWMutex() {
	mp = make(map[string]int)

	// 启动多个 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				key := "key"
				rwLock.RLock()
				_ = mp[key]
				rwLock.RUnlock()
			}
		}()
	}

	// 另一个 goroutine 进行写操作
	go func() {
		for i := 0; i < 100; i++ {
			key := "key"
			rwLock.Lock()
			mp[key]++
			rwLock.Unlock()
		}
	}()

}

// 关于 sync.Map 包，它是 Go 1.9 引入的一个新的并发安全的 map 实现，提供了读写锁的实现。
// 它的优点是代码简洁，读写锁的控制隐藏在底层实现中，不需要自己手动实现。
// 但是，sync.Map 的实现相对复杂，它使用了大量的内存预分配和锁分段等技巧来实现并发安全。
// 因此，在需要最高效率的场景，可能不适合使用 sync.Map 包
func mapSyncMap() {
	// 创建一个 sync.Map
	m := &sync.Map{}

	// 写入键值对
	m.Store("key1", "value1")

	// 读取键值对
	value, ok := m.Load("key1")
	if ok {
		fmt.Println("value:", value)
	}

	// 删除键值对
	m.Delete("key1")

	// 判断键是否存在
	_, ok = m.Load("key1")
	if !ok {
		fmt.Println("key1 does not exist")
	}

	// 遍历键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})
}

func mapSyncMapDemo() {
	// 创建一个 sync.Map
	var m sync.Map

	// 创建一个 wait group
	var wg sync.WaitGroup

	// 设置等待的协程数量为 1000
	wg.Add(1000)

	// 创建 1000 个协程同时写入键值对
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			m.Store(fmt.Sprintf("key-%d", idx), fmt.Sprintf("value-%d", idx))
			wg.Done()
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 读取所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})

}

func main() {
	mapSyncMutex()
	// mapRWMutex()
	// mapSyncMap()
	mapSyncMapDemo()
}
