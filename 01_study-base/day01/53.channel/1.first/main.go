package main

import (
	"fmt"
	"sync"
)

var a []int

// 需要指定通道中元素的类型
// 通道必须使用make函数初始化才能使用
var b chan int

var wg sync.WaitGroup

// channel 的入门
func main() {

	//buff()

	//buffMistake()

	noBuff()

}

// 带缓冲区的channel
func buff() {
	b = make(chan int, 1)
	// 写数据
	b <- 10
	fmt.Println("发送到通道中去了")
	// 读数据
	x := <-b
	fmt.Println("收到数据了:", x)
	close(b)
}

// 带缓冲区的channel的错误演示,死锁
// 指定了缓冲区的大小为1,当10写进入,就不能再写了,必须得读出来,才能继续写
func buffMistake() {
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

// 不带缓冲区的channel
func noBuff() {

	// 通道必须使用make函数初始化才能使用,不带缓冲区通道的初始化
	b = make(chan int)

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
