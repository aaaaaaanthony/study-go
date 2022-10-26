package main

import "fmt"

// 函数
func main() {

	sum := sum(10, 20)
	fmt.Println(sum)

	sum2 := sumForName(10, 20)
	fmt.Println(sum2)

	age, name := f5()
	fmt.Println(age, name)
}

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有参数也没有返回值
func f1() {
	fmt.Println("f2")
}

// 没有参数但有返回值的
// 返回值可以命名也可以不命名
func f3() int {
	return 3
}

func sumForName(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func f5() (int, string) {
	return 1, "antahony"
}

// 可变参数
// 必须放在参数的最后
// Go函数没有默认的参数
func f7(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 的类型是[]slice
}
