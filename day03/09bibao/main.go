package main

import "fmt"

// 毕包
// 闭包是一个函数,这个函数包含了他外部作用域的一个变量
func main() {
	ret := adder(100)
	res2 := ret(29)
	fmt.Println(res2)
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
