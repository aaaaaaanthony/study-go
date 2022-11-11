package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1.声明一个全局的变量
var wg sync.WaitGroup

// 使用sync.WaitGroup 去替换 线程睡眠的功能
func main() {

	for i := 0; i < 10; i++ {
		// 有一个线程就+1
		wg.Add(1)
		go f1(i)
	}
	// 等待
	wg.Wait()
	fmt.Println("完成了")
}

func f1(i int) {

	// 运行完就-1
	defer wg.Done()
	// 模拟业务处理
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}
