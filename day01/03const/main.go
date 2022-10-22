package main

import "fmt"

// 常量
const PI = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时,如果某一行声明后没有赋值,默认和上一行一样
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota // iota刚出现的时候默认是0
	a2        // const中iota每新增一行常量声明,就默认加+

)

const (
	b1 = iota
	b2
	_
	b3
)

const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {

	// 常量不能再修改
	// 	PI = 123

	fmt.Println(PI)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println()

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println()

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println()

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println()

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println()

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
}
