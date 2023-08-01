package main

/*
import "fmt"
import "time"
*/

import (
	"fmt"
	"time"
)

// main函数
func main() { // 函数的{ 一定和函数名在同一行
	// golang 中表达式 加 ;  和 不加 ; 都可以
	fmt.Println("Hello, world!")

	time.Sleep(1 * time.Second)
}
