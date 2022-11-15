package main

import (
	"fmt"
	"strings"
)

// 毕包
// 闭包是一个函数,这个函数包含了他外部作用域的一个变量
func main() {
	ret := adder(100)
	res2 := ret(29)
	fmt.Println(res2)

	fmt.Println("===============")
	jpg := myFunc(".jpg")
	txt := myFunc(".txt")

	fmt.Println(txt("test.txt"))
	fmt.Println(txt("test"))

	fmt.Println(jpg("test.jpg"))
	fmt.Println(jpg("test"))

	fmt.Println("===============")

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func myFunc(suffix string) func(string) string {

	return func(name string) string {

		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}

}

func calc(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		return base + 1
	}

	sub := func(i int) int {
		return base - 1
	}

	return add, sub
}
