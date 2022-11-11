package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 一个特殊情况的联系
// goroutine + 闭包
func main() {

	for i := 0; i < 1000; i++ {
		// 闭包
		// 这样打印的时候会有问题,会打印重复的值,因为当函数获取值的时候,i已经跑到很大值了
		//go func() {
		//	fmt.Println(i)
		//}()

		// 这样就是对的
		wg.Add(1)
		go func(i int) {
			wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("运行结束了")
}
