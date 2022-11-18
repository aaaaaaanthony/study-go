package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	for {
		var key string
		fmt.Println("输入信息")
		fmt.Scan(&key)
		fmt.Println(key)
	}
}
