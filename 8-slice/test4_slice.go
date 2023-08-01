package main

import "fmt"

/*
	切片容量的追加
*/
func main() {
	// 初始化 一个 slice 切片 长度 为 3 容量为5
	var number = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	// 添加一个元素 此时长度 为 4 容量为 5
	number = append(number, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	// 添加一个元素 此时长度 为 5 容量为 5
	number = append(number, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	// 添加一个元素 此时长度 为 6 容量为 10 切片 根据cap 容量 动态扩容： 容量* 2 = 10
	number = append(number, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	// 初始化 一个 slice 切片 长度 为 3 容量为3
	var number1 = make([]int, 3)

	// 添加一个元素 此时长度 为 4 容量为 6 切片 根据cap 容量 动态扩容： 容量* 2 = 6
	number1 = append(number1, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number1), cap(number1), number1)

}
