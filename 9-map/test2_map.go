package main

import "fmt"

func pritMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key =", key, "value =", value)
	}
	// 引用传递 修改原始数据
	myMap["three"] = "three"
}

/*
	map 增删改查
*/
func main() {
	myMap := make(map[string]string)
	// 增加
	myMap["one"] = "java"
	myMap["two"] = "python"
	myMap["three"] = "go"

	// 删除
	delete(myMap, "one")

	// 遍历
	pritMap(myMap)

	// 修改
	myMap["two"] = "abc"

	fmt.Println("==========")

	// 遍历
	pritMap(myMap)

	fmt.Println("==========")
	for key, value := range myMap {
		fmt.Println("key =", key, "value =", value)
	}
}
