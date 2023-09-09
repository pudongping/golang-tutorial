package main

import (
	"log"
)

// 这里会直接报错 `fatal error: concurrent map writes` 并直接终止程序 recover 无法捕获到错误
func main() {
	m := make(map[int]string)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("recover: %v", e)
				}
			}()

			m[i] = "Hello hhh"
		}()
	}

	log.Println("nice")
}
