package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("主线程结束运行")
	fmt.Println("主线程正在运行。。。")

	// 创建一个3个缓冲的channel
	c := make(chan int, 3)
	fmt.Printf("len(c) = %d, cap(c) = %d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子线程结束运行")
		fmt.Println("子线程正在运行。。。")

		// 当 channel 已经满 再向里面写数据 就会阻塞
		// 当 channel 为空 从里面取数据也会阻塞
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("len(c) = %d, cap(c) = %d\n", len(c), cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
}
