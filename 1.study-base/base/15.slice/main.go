package main

import "fmt"

// 切片
func main() {

	// 切片的定义
	var s1 []int // 定义一个存放int类型元素的切片
	var s2 []string
	fmt.Printf("s1的类型是%T,数据是:%v,s2的类型是%T,数据是:%v", s1, s1, s2, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"a", "b", "c"}
	fmt.Printf("s1的类型是%T,数据是:%v,s2的类型是%T,数据是:%v", s1, s1, s2, s2)
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 定义并且初始化
	var d = []string{"1", "2"}
	fmt.Println(d)

	// 切片的长度和容量
	// cap 数组的容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 数组转切片
	// 切片指向了一个底层的数组,
	// 切片的长度就是它的个数
	// 切片的容量是底层数组从切片的第一个元素到最后一个的数量
	a1 := [...]int{1, 2, 3, 4, 5}
	s3 := a1[0:4]
	fmt.Println(s3)
	s4 := a1[1:5]
	fmt.Println(s4)
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s4, s5, s6, s7)
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7))

	// 切片是引用类型,都指向底层的一个数组
	a1[4] = 1000
	fmt.Println(s7)
}
