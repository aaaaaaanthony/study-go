package main

import "fmt"

// make函数创建的切片 ,make只能用于 slice  map  chan   返回的是对应的这个三个类型的本身
// new很少用,一般用来给基本数据类型申请内存  stirng\int.... ,返回的是对应类型的指针
func main() {

	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v  len(s1)=%d   cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)

	s3[0] = 1000
	fmt.Println(s3, s4)

}
