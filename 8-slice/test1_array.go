package main

import "fmt"

/*
	数组
*/

func arrayPrint(array [4]int) {
	for index, value := range array {
		fmt.Println("index = ", index, "value = ", value)
	}
	// array 属于值拷贝，不影响myArray3的数据
	array[0] = 1000
}
func main() {
	// 固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	// 只能传递 [4]int 类型数据，属于值拷贝给形参array
	arrayPrint(myArray3)
}
