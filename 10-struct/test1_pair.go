package main

import "fmt"

/*
   pair
*/

type Reader interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("ReadBook")
}

func (this *Book) WriteBook() {
	fmt.Println("WriteBook")
}

func main() {
	// book<type:Book, value:book{}>
	book := &Book{}
	// r<type:Reader, value:>
	var r Reader
	// r<type:Book, value:book{}>
	r = book
	r.ReadBook()
	// w<type:Write, value:>
	var w Write
	// r<type:Book, value:book{}>
	w = r.(Write) // w interface Book 此时 w 和 r 类型是一致的
	w.WriteBook()
}
 