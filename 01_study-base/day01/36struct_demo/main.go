package main

import "fmt"

type person struct {
	name string
	age  int
}

// 结构体占用连续的空间
type x struct {
	a int8 // 8bit->1byte
	b int8
	c int8
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

	fmt.Println("=====================")

	// 结构体指针1
	// 变量p2是指针类型
	var p2 = new(person)
	fmt.Printf("%T new出来地址是:%p new出来的对象的变量地址是:%p\n", p2, p2, &p2)

	// 结构体指针2
	var p3 = person{
		name: "anthony",
		age:  12,
	}
	var p32 = &person{
		name: "anthony",
		age:  12,
	}
	fmt.Printf("%T %v %p\n", p3, p3, &p3)
	fmt.Printf("%T %p %p\n", p32, p32, &p32)

	// 列表的姓氏初始化
	var p4 = person{
		"anthony",
		12,
	}
	fmt.Printf("%T %v %p\n", p4, p4, &p4)

	fmt.Println("=====================")
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p \n", &(m.a))
	fmt.Printf("%p \n", &(m.b))
	fmt.Printf("%p \n", &(m.c))
}
