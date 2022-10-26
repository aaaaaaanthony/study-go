package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type yourInt = int

// 自定义类型和类型别名
func main() {

	var n myInt
	n = 100
	fmt.Printf("%T %v\n", n, n)

	var m yourInt
	m = 100
	fmt.Printf("%T %v\n", m, m)
}
