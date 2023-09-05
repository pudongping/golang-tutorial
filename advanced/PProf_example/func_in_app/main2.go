package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func init() {
	// 需要特别注意的是 runtime.SetMutexProfileFraction 语句，如果未来希望进行互斥锁的采集，
	// 那么需要通过调用该方法来设置采集频率，若不设置或没有设置大于 0 的数值，默认是不进行采集的
	runtime.SetMutexProfileFraction(1)
}

func main() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	_ = http.ListenAndServe(":6061", nil)
}
