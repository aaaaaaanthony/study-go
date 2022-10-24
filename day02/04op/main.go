package main

import "fmt"

func main() {

	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// 单独的语句,不能放在=的右边赋值
	a++
	b--
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符

	// 位运算符
	
}
