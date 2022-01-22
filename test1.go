package main

import (
	"fmt"
	"time"
)

func main()  {

	a := time.Now()
	fmt.Println(a.Unix())
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now().Unix())
	b := time.Since(a)
	fmt.Println(b)
}
