package main

import "fmt"

/*
	struct 类 方法
*/

// 类名、方法、属性 首字母大写 表示该属性是对外能够访问 否则的话只能够内部访问
type Hero struct {
	// 属性首字母大写 表示该属性是对外能够访问 否则的话只能够内部访问
	Name string
	Ad   int
	Leve int
}

/*
func (this Hero) SetName(name string) {
	this.Name = name
}

func (this Hero) GetName() string {
	return this.Name
}
*/
func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Leve = ", this.Leve)
}

func main() {
	hero := Hero{Name: "John", Ad: 1, Leve: 1}
	hero.Show()

	// 修改无效，需修改SetName方法 this *Hero
	hero.SetName("test")
	hero.Show()
}
