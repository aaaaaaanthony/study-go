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
	fmt.Println("θ΅°π±ζ­₯")
}

func (c cat) eat(str string) {
	fmt.Println("π±επ", str)
}

type chicken struct {
	feet int
}

func (c chicken) move() {
	fmt.Println("πε¨")
}

func (c chicken) eat(str string) {
	fmt.Println("πει₯²ζ", str)
}

func main() {

	var a1 animal
	bc := cat{
		name: "ζ·ζ°",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	a1.eat("ε°ι»θ±")

	ch := chicken{
		feet: 4,
	}
	a1 = ch
	a1.eat("θ‘εΎε€")
	fmt.Printf("%T\n", a1)

}
