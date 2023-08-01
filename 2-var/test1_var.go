package main

import "fmt"

/*
	四种变量
*/
// 声明全局变量 方法一、方法二、方法三可以
var h string = "100"
var j int = 200
var k = 300

// 用方法四声明全局变量 (报错)
// i := 400

func main() {
	// 方式一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明 一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s\ntype of bb = %T\n", bb, bb)

	// 方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s\ntype of cc = %T\n", cc, cc)

	// 方法四（最常用） 省去var关键字 直接自动匹配， := 只能用在函数体内
	e := 100
	fmt.Printf("type of e = %T\n", e)

	f := "abc"
	fmt.Printf("f = %s\ntype of f = %T\n", f, f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("===全局变量======")
	fmt.Println("h = ", h)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	// fmt.Println("i = ", i)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)

	// 多变量声明
	var kk, ll = 100, "aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
