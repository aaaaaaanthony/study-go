package main

import (
	"context"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func g(ctx context.Context) {
	value := ctx.Value("mykey")
	fmt.Println("从content获取数据:", value)
	defer wg.Done()
}

func main() {

	// 第一个参数:context.Background() 空的上下文
	// 第二个参数:key
	// 第三个参数:value
	ctx := context.WithValue(context.Background(), "mykey", "myvalue")

	fmt.Println("运行开始")
	wg.Add(1)
	go g(ctx)
	fmt.Println("运行完成")

	wg.Wait()

}
