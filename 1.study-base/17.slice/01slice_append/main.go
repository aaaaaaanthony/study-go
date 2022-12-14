package main

import "fmt"

func main() {

	s1 := []string{"北京", "上海", "深圳"}
	// s1[3] = "广州" 报错

	// append 必须要用原来的切片变量接收返回值
	s1 = append(s1, "广州")
	fmt.Println(s1)

	s2 := append(s1, "西安")
	fmt.Println(s2)

	s1 = append(s1, "佛山", "纽约")
	fmt.Println(s1)

	ss := []string{"伦敦", "法国"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Println(s1)
}
