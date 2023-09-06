package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("hello World!")
			time.Sleep(time.Millisecond * 500)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
