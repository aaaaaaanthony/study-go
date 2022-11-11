package main

import (
	"fmt"
	"strings"
)

func main() {

	path := "/User/Anthony"
	fmt.Println(path)

	// 多行的字符串,反引号
	s2 := `
锄禾日当午
	汗滴禾下锄
		谁知盘中餐
			粒粒皆辛苦
	`
	fmt.Println(s2)

	// 字符串的相关操作
	fmt.Println(len(s2))

	// 字符串拼接
	xing := "姓"
	ming := "名"
	fmt.Println(xing + ming)

	ss := xing + ming
	fmt.Println(ss)

	fmt.Sprintf("%s%s", xing, ming)

	// 分割
	hello := "Hello world"
	ret := strings.Split(hello, " ")
	fmt.Println(ret)
	fmt.Println(ret[0])
	fmt.Println(ret[1])

	// 包含
	fmt.Println(strings.Contains(hello, "o"))

	// 前缀
	fmt.Println(strings.HasPrefix(hello, "He"))
	// 后缀
	fmt.Println(strings.HasSuffix(hello, "world"))

	// 出现位置的索引
	s4 := "abcdef"
	fmt.Println(strings.Index(s4, "b"))
	fmt.Println(strings.LastIndex(s4, "e"))

	// Join操作
	fmt.Println(strings.Join(ret, "+"))

	//
}
