package main

import "fmt"

// 全局变量
var x = 100

// 变量的作用域
func main() {

	f1()
}

func f1() {
	fmt.Println(x)
}
