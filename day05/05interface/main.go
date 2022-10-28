package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println("走🐱步")
}

func (c cat) eat(str string) {
	fmt.Println("🐱吃🐟", str)
}

type chicken struct {
	feet int
}

func (c chicken) move() {
	fmt.Println("🐔动")
}

func (c chicken) eat(str string) {
	fmt.Println("🐔吃饲料", str)
}

func main() {

	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	a1.eat("小黄花")

	ch := chicken{
		feet: 4,
	}
	a1 = ch
	a1.eat("蔡徐坤")
	fmt.Printf("%T\n", a1)

}
