package main

import "fmt"

// map和slice组合
func main() {

	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)

	var m1 = make(map[int]string, 1)
	m1[1] = "1"
	s1[0] = m1

	s1[1] = make(map[int]string)
	s1[1][2] = "2"

	fmt.Println(s1)

	// 值为切片类型的map
	var m2 = make(map[string][]int, 10)
	m2["北京"] = []int{10, 20, 30}
	fmt.Println(m2)

}
