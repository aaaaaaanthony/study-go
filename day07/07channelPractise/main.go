package main

import (
	"fmt"
	"sync"
)

// channel联系
// 1.启动一个线程,生成100个数发送到ch1
// 2.启动一个线程,从ch1中取值,平方计算放到cha2中
func main() {

	a := make(chan int, 100)
	b := make(chan int, 100)

	go f1(a)
	go f2(a, b)
	wg.Add(2)

	for ret := range b {
		fmt.Println(ret)
	}
}

var wg sync.WaitGroup

func f1(ch1 chan int) {

	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)

}

func f2(ch1, ch2 chan int) {
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

// 单向通道
// cha1 只能放
// cha2 只能取
func f3(cha1 chan<- int, cha2 <-chan int) {
	cha1 <- 10
	//<-cha1

}
