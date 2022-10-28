package main

import "fmt"

type speaker interface {
	speak()
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (c dog) speak() {
	fmt.Println("旺旺旺")
}

func da(s speaker) {
	s.speak()
}

// 接口实例
func main() {

	var c1 cat
	var d1 dog

	da(c1)
	da(d1)

}
