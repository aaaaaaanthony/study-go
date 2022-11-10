package main

// import (
// 	"fmt"
// 	"os"
// )

// // var smr studentMgr = studentMgr{
// // 	allStudent: make(map[string]student, 48),
// // }

// var smr = studentMgr{
// 	allStudent: make(map[int]student, 48),
// }

// func showMenu() {

// 	fmt.Println(`
// 	欢迎使用系统
// 		1.查看所有学生
// 		2添加学生
// 		3.修改学生
// 		4.删除学生
// 		5.退出
// 	`)
// }

// // 学生管理系统
// func main() {

// 	for {
// 		showMenu()
// 		fmt.Println("请输入序号:")
// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			smr.all()
// 		case 2:
// 			smr.add()
// 		case 3:
// 			smr.update()
// 		case 4:
// 			smr.del()
// 		case 5:
// 			os.Exit(1)
// 		default:
// 			fmt.Println("输入正确的序号")
// 		}
// 	}
// }
