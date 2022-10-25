package main

import "fmt"

func main() {

	// 求和
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0

	for _, one := range a1 {
		sum = sum + one
	}
	fmt.Println(sum)
}
