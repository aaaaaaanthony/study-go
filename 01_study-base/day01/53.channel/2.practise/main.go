package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// channel联系
// 1.启动一个线程,生成100个数发送到ch1
// 2.启动一个线程,从ch1中取值,平方计算放到cha2中
func main() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go receiveFromCh1(ch1)
	go receiveFromCh2(ch1, ch2)
	wg.Add(2)

	for ret := range ch2 {
		fmt.Println(ret)
	}
}

func receiveFromCh1(ch1 chan int) {

	defer wg.Done()
	for i := 0; i < 100; i++ {
		// 写数据到ch1
		ch1 <- i
	}
	// 这就关闭了channel,还是可以从channel接收到数据的
	close(ch1)

}

func receiveFromCh2(ch1, ch2 chan int) {
	for {
		// 从ch1接收消息
		x, ok := <-ch1
		if !ok {
			break
		}
		// 计算好之后 写进ch2
		ch2 <- x * x
	}
	close(ch2)
}
