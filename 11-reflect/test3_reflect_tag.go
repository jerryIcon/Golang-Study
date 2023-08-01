package main

import (
	"fmt"
	"reflect"
)

/*
tag
*/
type Book struct {
	Name string `info:"name" doc:"我的名字"`
	Auth string `info:"sex"`
}

func findTag(arg interface{}) {
	e := reflect.TypeOf(arg).Elem()
	for i := 0; i < e.NumField(); i++ {
		info := e.Field(i).Tag.Get("info")
		doc := e.Field(i).Tag.Get("doc")
		fmt.Println("info = ", info, " doc = ", doc)
	}
}

func main() {
	b := Book{"jer", "asc"}
	findTag(&b)
}
