package main

import (
	"fmt"
	"time"
)

func TrackTime(position string) func() {
	pre := time.Now()
	return func() {
		elapsed := time.Since(pre)
		fmt.Printf("[%s] elapsed: %s\n", position, elapsed)
	}
}

func main() {
	defer TrackTime("time_helper main")()

	fmt.Println("hello Go!")
	time.Sleep(1 * time.Second)
}
