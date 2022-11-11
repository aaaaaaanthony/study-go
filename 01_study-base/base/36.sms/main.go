package main

import (
	"fmt"
	"os"
)

var (
	// 声明变量,还没有创建
	all map[int]*student
)

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}

}

func list() {
	for k, v := range all {
		fmt.Printf("学号:%d,姓名:%s\n", k, v.name)
	}

}

func addStudent() {

	var id int
	var name string

	fmt.Println("请输入学生的学号")
	fmt.Scanln(&id)

	fmt.Println("请输入学生的姓名")
	fmt.Scanln(&name)

	// 创建一个学生
	s := newStudent(id, name)

	// 添加
	all[id] = s
}

func del() {

	var id int
	fmt.Println("输入要删除学生的学号!")
	fmt.Scanln(&id)

	delete(all, id)

}

// 函数版学生管理系统
func main() {

	// 开辟内存空间
	all = make(map[int]*student, 48)

	for {
		fmt.Println(`
		欢迎光临学生管理系统!
			1.查看所有学生
			2.新增学生
			3.删除学生
			4.退出
		`)
		fmt.Println("请输入你要干啥")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你学则了%d这个选项!\n", choice)

		switch choice {
		case 1:
			list()
		case 2:
			addStudent()
		case 3:
			del()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("重新输入")
		}
	}

}
