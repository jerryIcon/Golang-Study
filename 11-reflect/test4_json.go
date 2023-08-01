package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	m := Movie{"title", 2023, 10000, []string{"xing", "jie"}}
	// struct 转 json
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换错误", err)
		return
	}
	fmt.Println("jsonStr = ", string(jsonStr))
	// json 转 struct
	var m1 Movie
	err = json.Unmarshal(jsonStr, &m1)
	fmt.Printf("%T : %v\n", m1, m1)
}
