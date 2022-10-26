package main

import "fmt"

// 指针
func main() {

	// 1.取地址
	n := 18
	p := &n
	fmt.Printf("%T,%p , %p \n", p, p, &n)

	// 2. 根据地址取值
	m := *p
	fmt.Printf("%T,%v\n", m, m)

	// 声明一个int类型的指针
	var a1 *int
	//*a1 = 100
	//fmt.Println(*a1)
	fmt.Println(a1)
	var a2 = new(int) // new申请内存地址
	*a2 = 100
	fmt.Println(*a2)
}
