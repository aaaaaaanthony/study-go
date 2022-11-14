package main

import "fmt"

func main() {

	age := 19
	if age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("没有成年")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	// 作用域
	// age2 在if的条件判断作用域
	if age2 := 19; age2 > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("没有成年")
	}
	// 这里找不到age2
	// fmt.Println(age2)

}
