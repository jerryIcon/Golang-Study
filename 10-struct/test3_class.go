package main

import "fmt"

/*
   继承，重载
*/
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat()...")
}

func (this *Human) Run() {
	fmt.Println("Human Run()...")
}

type SuperMan struct {
	Human // SuperMan 继承 Human
	Leve  int
}

// 子类重载父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("Leve = ", this.Leve)
}

func main() {
	s := SuperMan{Human{"jack", "girl"}, 100}
	s.Eat() // 调用子类重载方法
	s.Run() // 父类方法

	fmt.Println("=================")
	var s1 SuperMan
	s1.name = "s1"
	s1.sex = "boy"
	s1.Leve = 2000
	s1.Print()
}
