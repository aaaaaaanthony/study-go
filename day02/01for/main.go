package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出for循环
		}
		fmt.Println(i)
	}

	fmt.Println("=================")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("=================")
	for i := 0; i < 10; {
		if i == 5 {
			continue
		}
		fmt.Println(i)
		i++
	}
}
