package main

import (
	"expvar"
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

type upTimeVar struct {
	value time.Time
}

func (v *upTimeVar) Set(date time.Time) {
	v.value = date
}

func (v *upTimeVar) Add(duration time.Duration) {
	v.value = v.value.Add(duration)
}

func (v *upTimeVar) String() string {
	return fmt.Sprintf("\"%s\"", v.value.Format(time.RFC3339))
}

var (
	appleCounter      *expvar.Int
	GOMAXPROCSMetrics *expvar.Int
	upTimeMetrics     *upTimeVar
)

func init() {
	appleCounter = expvar.NewInt("apple")
	GOMAXPROCSMetrics = expvar.NewInt("GOMAXPROCS")
	GOMAXPROCSMetrics.Set(int64(runtime.NumCPU()))

	upTimeMetrics = &upTimeVar{value: time.Now().Local()}
	expvar.Publish("uptime", upTimeMetrics) // 发布自定义类型指标
}

func main() {
	log.Println("Server at http://127.0.0.1:6060")

	// 获取指定指标
	expvarFunc := expvar.Get("memstats").(expvar.Func)
	memstats := expvarFunc().(runtime.MemStats)
	fmt.Printf("memstats.GCSys: %d \n", memstats.GCSys)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		appleCounter.Add(1)
		fmt.Println(r.URL)
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":6060", http.DefaultServeMux)
}
