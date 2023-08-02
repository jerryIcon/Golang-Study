package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			if i == 3 {
				// close 关闭一个channel通道
				close(c)
				break
			}
			c <- i
		}
	}()

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finished")
}
