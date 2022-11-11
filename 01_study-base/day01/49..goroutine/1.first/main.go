package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("goroutine")
}

// goroutine 入门使用
func main() {

	go hello()
	fmt.Println("main")
	// 需要等待,如果不等待,肯能看不到hello打印,因为hello还没有打印的时候,main方法已经运行完了
	time.Sleep(time.Second)
}
