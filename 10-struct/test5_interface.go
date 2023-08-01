package main

import "fmt"

/*
   接口 interface{} 万能数据类型
*/

func myFunc(arg interface{}) {
	fmt.Println("my Func() ....")
	fmt.Println(arg)

	// 给 interface{} 断言 "类型断言" 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg 不是 string 类型")
	} else {
		fmt.Println("arg 是 string 类型")
		fmt.Printf("type %T\n", value)
	}
}

type Book struct {
	Name string
}

func main() {
	book := Book{"jack"}
	myFunc(book)

	myFunc("asd")
	myFunc(123)
	myFunc(123.567)
}
