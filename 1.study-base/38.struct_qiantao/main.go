package main

import "fmt"

type person struct {
	name string
	age  int
	addr address

	boss      // 匿名嵌套结构体
	workPlace // 匿名嵌套结构体 ,这两个匿名嵌套结构体都有name
}

type company struct {
	name string
	addr address
}

type workPlace struct {
	province string
	city     string
	name     string
}

type address struct {
	province string
	city     string
}

type boss struct {
	name    string
	bossAge int
}

// 结构体嵌套
func main() {

	p1 := person{
		name: "anthony",
		age:  18,
		addr: address{
			province: "广东",
			city:     "广州",
		},
		boss: boss{
			name:    "boss",
			bossAge: 99,
		},
		workPlace: workPlace{
			province: "佛山",
			city:     "千灯湖",
			name:     "工作地点",
		},
	}

	fmt.Println(p1)
	fmt.Println(p1.addr.province)
	fmt.Println("========================")

	// 匿名嵌套结构体的打印
	fmt.Println(p1.name) // 本来是想打印boss的name的,  原理是先在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找字段
	fmt.Println(p1.boss.name)
	fmt.Println(p1.boss.bossAge)
	fmt.Println("========================")

	//  字段冲突,就要完全指定字段
	fmt.Println(p1.name)
	fmt.Println(p1.workPlace.name)
	fmt.Println(p1.boss.name)

}
