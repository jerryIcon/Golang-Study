package main

import "fmt"

/*
	切片声明方式
*/

func main() {
	// 声明切片四种方式

	// 1. 声明slice 是一个切片，但是没有给slice 分配空间
	var slice []int

	// 2.声明slice是一个切片 同时给slice分配3个空间 默认值是 0
	// var slice []int = make([]int, 3)

	// 3.声明slice 是一个切片，并且初始化默认值为 1,3,4 长度是3
	// slice := []int{1, 3, 4}

	// 4.声明slice 是一个切片， 同时给slice分配3个空间 默认值是 0
	// slice := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice), slice)

	if slice == nil {
		fmt.Println("slice 是一个空切片")
	} else {
		fmt.Println("slice 不是一个空切片")
	}

}
