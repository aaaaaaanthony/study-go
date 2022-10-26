package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {

	// 声明一个person类型的变量
	var anthony person
	// 赋值
	anthony.name = "anthony"
	anthony.age = 11
	anthony.gender = "男"
	var hobby = [...]string{"足球", "篮球"}
	anthony.hobby = hobby[:]
	fmt.Printf("%T,%v\n", anthony, anthony)

	// 访问变量的字段
	fmt.Println(anthony.name)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "anthony"
	s.age = 12
	fmt.Printf("%T,%v\n", s, s)
}
