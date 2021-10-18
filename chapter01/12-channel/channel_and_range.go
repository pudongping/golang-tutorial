/*
channel 和 range
*/
package main

import "fmt"

func main() {

	c := make(chan int) // 创建一个无缓冲的 channel

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close 可以关闭一个 channel
		close(c)
	}()

	/*
		for {
			// 变量 ok 如果为 true 表示 channel 没有关闭，如果为 false 表示 channel 已经关闭
			// 这里的 data 和 ok 其实在 if 判断语句中属于局部变量
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	// 以下这种写法和以上的写法其实是一致的
	// 可以使用 range 来迭代不断操作 channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished ……")
}

/*
0
1
2
3
4
Main Finished ……
*/
