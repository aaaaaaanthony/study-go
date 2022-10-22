package main

import "fmt"

// 声明变量,Go语言中推荐使用驼峰式命名 userName
// var name string
// var age int
// var isOk bool

// 漂亮声明
var (
	name2 string // ""
	age2  int    // 0
	isOk2 bool   // false
)

func main() {

	name2 = "anthony"
	age2 = 18
	isOk2 = true

	// GO语言中声明变量必须使用,不使用就编译不过去
	fmt.Print(isOk2) // 在终端中输入要打印的内容,不会有换行符

	fmt.Printf("name2:%s\n", name2) // %s:占位符,使用那么这个变量的值去替换占位符
	fmt.Println(age2)               // 打印玩指定的内容之后会自动在后面加一个换行符

	// 生命变量同时赋值
	var s1 string = "anthony2"
	fmt.Println(s1)

	// 自动推导
	var s2 = "anthony3"
	fmt.Println(s2)

	// 简短变量生命,只能在函数里面使用
	s3 := "anthony4"
	fmt.Println(s3)

	// 匿名变量不占用命名空间,也不会分配内存,所以匿名变量之间不存在重复声明
	// 下划线_ 就是匿名变量
	// s1 := "anthony5" 在同一个作用域里不能重复声明
	a, _ := noSign()
	fmt.Println(a)

}

func noSign() (int, int) {
	return 5, 6
}
