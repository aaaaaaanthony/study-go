package main

import "fmt"

// 函数类型
func main() {

	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
}

func f1() {
	fmt.Println("你好")
}

func f2() (x int) {
	fmt.Println("你好")
	return 5
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	a := x()
	fmt.Println(a)
}

// 函数可以作为返回值

func ff(a, b int) int {
	return a + b
}

func f5(x func() int) func(int, int) int {

	return ff
}
