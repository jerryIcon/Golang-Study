package main

import "fmt"

/*
	struct结构体 相当于类
*/

// 声明一种行数据类型 myint 是int的一个别名
type myint int

// 定义一个结构体 Book
type Book struct {
	title string
	auth  string
}

// 值传递
func changBook1(book Book) {
	book.title = "no book"
}

func changBook2(book *Book) {
	book.title = "no book"
}

func main() {
	var a myint = 10
	fmt.Printf("a = %v\n", a)

	var book Book
	book.title = "Book"
	book.auth = "zhang3"
	changBook1(book)
	// 修改 book 修改失败
	fmt.Printf("book = %v\n", book)

	changBook2(&book)
	// 修改 book 修改成功
	fmt.Printf("book = %v\n", book)
}
