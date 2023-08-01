package main

import "fmt"

func deferFunc() int {
	fmt.Println("deferfunc")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer关键字 在函数体结束之前执行 defer 以压栈形式，先进后出
	defer fmt.Println("main end1") // 后出栈
	defer fmt.Println("main end2") // 先出栈

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
	fmt.Println("======================")
	returnAndDefer()   // return 先执行 defer后执行
}
