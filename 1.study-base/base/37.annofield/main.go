package main

import "fmt"

type person struct {
	string
	int
}

func main() {

	p1 := person{
		"anthony",
		12,
	}

	fmt.Println(p1.int)
	fmt.Println(p1.string)
}
