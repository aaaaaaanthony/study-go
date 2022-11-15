package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
)

// flag库
func init() {

	// go build
	// 用不用等号效果是一样的
	// $ ./1.flag -intFlag=12 -boolFlag=true -stringFlag=abc
	// $ ./1.flag -intFlag 12 -boolFlag true -stringFlag abc

	// $ ./1.flag -help

	//  方法里的第二个参数,是指定名字
	flag.IntVar(&intFlag, "intFlag", 0, "整数的值")
	flag.BoolVar(&boolFlag, "boolFlag", false, "布尔类型的值")
	flag.StringVar(&stringFlag, "stringFlag", "default", "字符串类型的值")

}

func main() {

	flag.Parse()

	// 第一种方法
	fmt.Println("int:", intFlag)
	fmt.Println("boolflag:", boolFlag)
	fmt.Println("stringflag:", stringFlag)

}
