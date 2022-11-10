package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exitChan = make(chan bool, 1)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)

	go f(ctx)
	time.Sleep(time.Second * 5)

	// 通知goroutine结束
	cancel()

	wg.Wait()
}

func f(ctx context.Context) {

	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("111")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("222")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}
