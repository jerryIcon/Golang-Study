package main

import (
	"fmt"
	"reflect"
)

/*
reflect 反射
*/
type Book struct {
	Name string
	Auth string
}

func (this Book) Called() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	// 通过反射获取类型
	inputType := reflect.TypeOf(input)
	fmt.Println("type: ", inputType)

	// 通过反射获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value: ", inputValue)

	// 通过反射获取Filed字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("type : ", field, " value : ", value)
	}

	// 通过放射获取method
	fmt.Println(inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : % v\n", m.Name, m.Type)
	}
}

func main() {
	book := Book{"book", "auther"}
	DoFiledAndMethod(&book)
}
