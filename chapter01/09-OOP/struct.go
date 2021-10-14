/*
定义一个结构体
 */
package main

import "fmt"

// 定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book) {
	// 传递一个 book 的副本
	book.auth = "harry"
}

func changeBook2(book *Book)  {
	// 指针传递
	book.auth = "Mark"
}

func main()  {

	var book1 Book
	book1.title = "Study-Golang"
	book1.auth = "Alex"
	// {Study-Golang Alex}
	fmt.Printf("%v\n", book1)

	changeBook(book1)
	// {Study-Golang Alex}
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	// {Study-Golang Mark}
	fmt.Printf("%v\n", book1)

}