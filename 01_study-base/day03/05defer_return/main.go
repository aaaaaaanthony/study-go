package main

import "fmt"

func main() {

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x,不是返回值
	}()
	return x
}

func f2() (x int) {

	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return 5 // 返回值=y=x=5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值=y=x=5
}
