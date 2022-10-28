package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) walk() {
	fmt.Println(p.name + "在走路")
}

// 值接收者,传拷贝
func (p person) guonian() {
	p.age = p.age + 1
}

// 指针接受者,传内存地址
func (p *person) guonian2() {
	p.age = p.age + 1
}

// 构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}

}
func main() {
	p1 := newPerson("anthony", 12)
	p2 := newPerson("anthony2", 13)
	fmt.Println(p1, p2)

	p1.walk()
	p2.walk()

	fmt.Printf("p1的年纪是:%d\n", p1.age)
	p1.guonian()
	fmt.Printf("p1的年纪是:%d\n", p1.age)
	p1.guonian2()
	fmt.Printf("p1的年纪是:%d\n", p1.age)
	p1.guonian()
	fmt.Printf("p1的年纪是:%d\n", p1.age)
}
