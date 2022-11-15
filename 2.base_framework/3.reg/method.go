package main

import (
	"fmt"
	"regexp"
	"strings"
)

var text = "https://baidu.com/123.png"

// Test1 找到有没有符合正则表达式的字符串
func Test1() bool {
	matched, err := regexp.MatchString("baidu.com", text)
	if err != nil {
		fmt.Println(err)
		panic("报错了")
	}
	return matched
}

func test3() {
	text := "https://aws.com/123.png"
	replace := strings.Replace(text, "aws", "image.mmzcg.com", -1)
	fmt.Println(replace)
}

func test2() {
	text := "https://aws.com/123.png"

	// 查找 Abc 或 a7c，替换为 Abccc a7ccc
	reg := regexp.MustCompile(`aws.com`)
	fmt.Println(reg)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, `ccc`)) //"Abccc a7ccc MFC 8ca. 你好！ Golang/"

}

func test() {
	buf := "abc"

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`s`)
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	replace := strings.Replace(buf, "a", "xxx", -1)
	fmt.Println(replace)
}
