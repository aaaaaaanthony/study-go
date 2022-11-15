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

	// 数组的遍历
	var a = [...]string{"a", "b", "c"}

	for i := 0; i < len(a); i++ { // 方法1
		fmt.Println(a[i])
	}

	for index, valuea := range a { // 方法2
		fmt.Println(index, valuea)
	}

	// 多维数组
	var a11 [3][2]int
	fmt.Println(a11)
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{1, 2},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 4}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

}
