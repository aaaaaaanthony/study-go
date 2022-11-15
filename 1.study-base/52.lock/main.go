package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

// 互斥锁,一个时间 只能读写一个资源
func main() {

	// 互斥锁
	test1()

	// 读写互斥锁
	test2()

}

func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg.Done()

}

func read() {
	defer wg.Done()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func test2() {

	now := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))

}

func test1() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
