package main

import (
	"time"
)

func main() {
	// 用 go 创建承载一个形参为空 返回值为空的一个匿名函数
	/*
		go func() {
			defer fmt.Println("A defer")
			func() {
				defer fmt.Println("B defer")
				runtime.Goexit()		// 终止当前的 goroutine
				fmt.Println("B")
			}()
			fmt.Println("A")
		}()
	*/

	// go func(a int, b int) bool {
	// 	fmt.Println("a = ", a, " b = ", b)
	// 	return true
	// }(1, 2)

	for {
		time.Sleep(1 * time.Second)
	}
}
