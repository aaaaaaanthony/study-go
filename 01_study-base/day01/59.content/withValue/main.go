package main

import (
	"fmt"
	"time"
)

func g2() {

}

func g() {
	fmt.Println("你是猪")
}

func main() {

	// 第一个参数:context.Background() 空的上下文
	// 第二个参数:key
	// 第三个参数:value
	//ctx := context.WithValue(context.Background(), "mykey", "myvalue")

	fmt.Println("运行开始")
	go g()
	fmt.Println("运行完成")

	time.Sleep(time.Second)

}
