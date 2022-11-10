package main

import (
	"fmt"
	"sync"
)

var a []int

// 需要指定通道中元素的类型
var b chan int

var wg sync.WaitGroup

// channel
func main() {
	// nil
	//Nobuff()

	//buff()

}

func buff2() {
	b = make(chan int, 1)
	wg.Add(1)
	b <- 10

	// 卡住了
	b <- 20
	fmt.Println("发送到通道中去了")
	x := <-b
	fmt.Println("收到数据了:", x)
	close(b)
}

func buff() {
	b = make(chan int, 1)
	wg.Add(1)
	b <- 10
	fmt.Println("发送到通道中去了")
	x := <-b
	fmt.Println("收到数据了:", x)
	close(b)
}

func Nobuff() {
	fmt.Println(b)
	// 通道必须使用make函数初始化才能使用,不带缓冲区通道的初始化
	b = make(chan int)
	// 带缓冲区的
	//b = make(chan int, 16)
	// 发送  chi1 <- 1
	// 接收  <- chi1

	wg.Add(1)
	go func() {
		x := <-b
		fmt.Println("从通道中收到消息了,数值是:", x)
		wg.Done()
	}()
	b <- 10
	fmt.Println("10发送到通道中了")
	wg.Wait()
	close(b)
}
