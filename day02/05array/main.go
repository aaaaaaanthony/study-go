package main

import "fmt"

// 数组
func main() {

	// 数组的声明
	// 必须指定存放元素的类型和长度
	// 数组的长度
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("%T,%T \n", a1, a2)

	// 数组的初始化
	// 如果不初始化,默认元素的零值
	fmt.Println(a1, a2)
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方法2
	a10 := [10]int{1, 2, 34, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	// 3.初始化方法3
	a100 := [...]int{1, 2, 4}
	fmt.Println(a100)
	// 4.初始化方法4
	a3 := [5]int{1, 2}
	fmt.Println(a3)
	// 5. 初始化方法5
	a4 := [5]int{0: 1, 3: 4}
	fmt.Println(a4)

}
