package main

import "fmt"

func main() {

	funcA()
	funcB()
	funcC()
}

func funcA() {

	fmt.Println("a")
}

func funcB() {

	// 刚打开数据库链接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库")
	}()

	// 程序崩溃退出
	panic("出现了严重的错误")
	// fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}
