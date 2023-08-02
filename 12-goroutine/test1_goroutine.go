package main

import (
	"fmt"
	"time"
)

/*
goroutine 协程
*/

func newTastk() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// go程
	go newTastk()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
