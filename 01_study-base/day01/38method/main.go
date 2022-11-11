package main

import "fmt"

// 报错
// func (i int) hello() {
// fmt.Print(i)
// }

type myInt int

func (i myInt) hello() {
	fmt.Println(i)
}

// 给自定义类型加方法
// 只能给自己包里的变量添加方法
func main() {

	my := myInt(100)
	fmt.Println(my)

}
