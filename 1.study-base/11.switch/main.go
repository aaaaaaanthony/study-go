package main

import "fmt"

func main() {

	switch n := 5; n {
	case 1:
		fmt.Println(1)
		break
	case 2:
		fmt.Println(2)
		break
	default:
		fmt.Println("没有找到")
	}
}
