package main

import "fmt"

type person struct {
	name string
	age  int
}

// 值类型,就是拷贝
func f(x person) {
	fmt.Println(x.name, x.age)
	x.name = "anthjony2"
}

func f2(x *person) {
	var my = *x
	fmt.Println(my)
	fmt.Println((*x).name, (*x).age)
	(*x).name = "anthony 指针传递修改值"

	// 语法糖
	fmt.Println(x.name)
	x.name = "anthony 指针传递修改值22222"
}

func main() {
	var p person
	p.name = "anthony"
	p.age = 12

	fmt.Println("=====================")
	f(p)
	fmt.Println(p.name)

	fmt.Println("=====================")
	f2(&p)
	fmt.Println(p.name)

	// 结构体指针1
	// 变量p2是指针类型
	var p2 = new(person)
	fmt.Printf("%T %p %p\n", p2, p2, &p2)

	// 结构体指针2
	var p3 = person{
		name: "anthony",
		age:  12,
	}
	fmt.Printf("%T %v %p\n", p3, p3, &p3)
}
