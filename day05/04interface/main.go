package main

import "fmt"

type car interface {
	run()
}

type xiaomi struct {
	brand string
}

func (x *xiaomi) run() {
	fmt.Println(x.brand)
}

type hengda struct {
	brand string
}

func (h *hengda) run() {
	fmt.Println(h.brand)
}

func drive(c car) {
	c.run()
}

func main() {

	var v1 = xiaomi{
		brand: "小米",
	}
	v1.run()

	var h1 = hengda{
		brand: "hengda",
	}
	v1.run()
	h1.run()

	var v2 xiaomi
	fmt.Printf("%p\n", &v2)

}
