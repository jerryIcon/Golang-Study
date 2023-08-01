package main

import (
	// "_ lib1"  		// _ 给当前包起匿名 无法使用当前包方法 但会执行当前包的内部init()方法
	// mylib1 lib1 	// 别名
	// . lib1				// 当前包的全部方法 导入当前包作用 可以直接使用 Lib1Test() 不需要lib1.Lib1Test()
	"lib2"
)

func main() {
	// mylib1.Lib1Test() // 别名调用
	lib1.Lib1Test()
	lib2.Lib12est()
}
