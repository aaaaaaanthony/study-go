package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func main() {

	for i := 0; i < 1000; i++ {
		// 闭包
		// 这样打印的时候会有问题,因为当函数获取值的时候,i已经跑到很大值了
		//go func() {
		//	fmt.Println(i)
		//}()

		// 这样就是对的
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}
