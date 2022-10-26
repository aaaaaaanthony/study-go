package main

import "fmt"

//defer 语句
func main() {

	deferDemo()
}

func deferDemo() {

	fmt.Println("======")
	// defer 吧后面的语句延迟到函数即将返回时再执行
	// 多个defer ,执行顺便,从后向前
	defer fmt.Println("11111")
	defer fmt.Println("22222")
	fmt.Println("======")
}
