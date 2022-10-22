package main

import "fmt"

// 整形
func main() {

	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%T\n", i4)
}
