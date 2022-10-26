package main

import "fmt"

func main() {

	a1 := []int{1, 3, 5}
	// 赋值
	a2 := a1

	// copy,完全拷贝,不管引用
	var a3 = make([]int, 3, 3)
	copy(a3, a1)

	fmt.Println(a1, a2, a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3)

	fmt.Println("=============================")

	// TODO 有问题
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1[0] = 200
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1[1] = 300
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1[2] = 300
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	fmt.Println("=============================")

	// x1 := [...]int{1, 3, 5}
	// s1 := x1[:]
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1, len(x1), cap(x1))

	// x1[0] = 200
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1, len(x1), cap(x1))

	// s1[0] = 300
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1, len(x1), cap(x1))
}
