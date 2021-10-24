/*
保证局部变量的安全性
*/
package main

import "fmt"

func main() {

	var j int = 1
	f := func() {
		var i int = 1
		fmt.Printf("i, j: %d, %d\n", i, j)
	}

	f()
	j += 2
	f()
	/*
		i, j: 1, 1
		i, j: 1, 3
	*/

}
