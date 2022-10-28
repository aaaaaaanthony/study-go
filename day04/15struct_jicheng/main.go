package main

import "fmt"

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%s在动\n", a.name)
}

type dog struct {
	feet int
	animal
}

func (a *dog) wang() {
	fmt.Printf("%s在叫:旺旺旺\n", a.name)
}

// 结构体模拟实现其它语言中的集成
func main() {

	d1 := dog{
		feet: 4,
		animal: animal{
			name: "小狗",
		},
	}

	d1.wang()
	d1.move()
}
