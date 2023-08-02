package main

import "fmt"

func main() {
	defer fmt.Println("gomain 结束")

	// 创建一个 channel
	c := make(chan int) // 无缓冲

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")

		c <- 666 // 传递int数据
	}()
	fmt.Println("gomain 正在运行")
	// 接收数据
	num := <-c
	fmt.Println("num = ", num)
}
