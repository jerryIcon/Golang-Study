package main

import "fmt"

func swap(p1 *int, p2 *int) {
	var temp int
	temp = *p1
	*p1 = *p2
	*p2 = temp
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	swap(&a, &b)
	fmt.Println("====================")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("====================")

	var p *int
	p = &a

	fmt.Println("&a = ", &a)
	fmt.Println("p = ", p)

	var pp **int
	pp = &p
	fmt.Println("&p = ", &p)
	fmt.Println("pp = ", pp)

}
