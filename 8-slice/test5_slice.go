package main

import "fmt"

/*
	切片截取
*/
func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	// [0,2) 左包右闭
	slice1 := slice[0:2] // 0, 1

	fmt.Println(slice1)

	// slice 和 slice1 指向同一个地址
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice)

	// 可以通过copy 来进行深拷贝
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Println(slice2)
	slice2[0] = 0
	fmt.Println(slice2)
}
