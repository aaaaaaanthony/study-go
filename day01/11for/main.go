package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}

	fmt.Println("====================")

	// 省略1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("====================")

	// 省略2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("====================")

	// 死循环
	// for {
	// fmt.Println(".......")
	// }

	// for range
	s := "Hello沙河"
	for i, v := range s {
		fmt.Println(i, string(v))
	}

}
