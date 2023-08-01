package main

import "fmt"

/*
   接口
*/

// 定义一个接口
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("car sleep()")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleep()")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "dog"
}

func showAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	animal := Cat{"Green"}
	animal.Sleep()
	showAnimal(&animal)

	animal1 := Dog{"red"}
	showAnimal(&animal1)
}
