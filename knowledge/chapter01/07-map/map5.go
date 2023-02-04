package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// unsafeMapDemo()
	safeMapDemo()

}

// 非线程安全的 map 示例
func unsafeMapDemo() {
	normalMap := make(map[int]string)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			normalMap[a] = fmt.Sprintf("我是a#{%d}", a)
			time.Sleep(time.Millisecond * 100)
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(normalMap[a])
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("unsafeMapDemo")
}

// 线程安全的 map 示例
func safeMapDemo() {
	m := New()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			m.Set(a, fmt.Sprintf("我是a#{%d}", a))
			time.Sleep(time.Millisecond * 100)
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(m.Get(a))
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("safeMapDemo")
}

// 剩下的写一个线程安全的 map
type SafeMap interface {
	Get(k interface{}) interface{}
	Set(k interface{}, v interface{})
}

type SafeMapImpl struct {
	SafeMap
	sync.RWMutex
	data map[interface{}]interface{}
}

func New() SafeMap {
	return &SafeMapImpl{
		data: make(map[interface{}]interface{}),
	}
}

func (s *SafeMapImpl) Get(k interface{}) interface{} {
	s.Lock()
	tmp := s.data[k]
	s.Unlock()
	return tmp
}

func (s *SafeMapImpl) Set(k interface{}, v interface{}) {
	s.Lock()
	s.data[k] = v
	s.Unlock()
}
