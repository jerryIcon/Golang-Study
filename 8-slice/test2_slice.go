package main

import "fmt"

/*
	切片
*/

// 动态数组  引用传递
func printArray(array []int) {
	// _ 表示匿名
	for _, value := range array {
		fmt.Println("value = ", value)
	}
	array[0] = 10000
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组 切片 slice
	fmt.Printf("myArray types %T\n", myArray)
	fmt.Println("===================")
	printArray(myArray)
	fmt.Println("===================")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
