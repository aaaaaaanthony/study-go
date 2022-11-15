package main

import "fmt"

// select学习
func main() {

	test()
}

// 一次select,只运行一次,这里的for循环,可以让select运行10次
func test() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}
}
