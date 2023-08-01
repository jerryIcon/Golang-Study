package main

import "fmt"

/*
	map三种声明方式
*/
func main() {
	// 1.声明空map
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空")
	}

	// 2.声明map 并开辟空间
	var myMap2 = make(map[int]string, 3)
	myMap2[1] = "java"
	myMap2[2] = "java"
	myMap2[3] = "java"
	fmt.Println(myMap2)

	// 3.声明map 并赋值
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "asd",
		"three": "cdz",
	}
	fmt.Println(myMap3)
}
