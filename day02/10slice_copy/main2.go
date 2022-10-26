package main

import "fmt"

func main() {

	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	s1[0] = 100
	// fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	// s1[1] = 300
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1, len(x1), cap(x1))

	// s1[2] = 300
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1, len(x1), cap(x1))

}
