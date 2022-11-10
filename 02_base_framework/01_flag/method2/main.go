package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    *int
	boolFlag   *bool
	stringFlag *string
)

// flag库
func init() {

	// 第一种定义的方法
	//method1()

	intFlag = flag.Int("intflag", 0, "int flag value")
	boolFlag = flag.Bool("boolflag", false, "bool flag value")
	stringFlag = flag.String("stringflag", "default", "string flag value")

}

func main() {

	flag.Parse()

	fmt.Println("int:", *intFlag)
	fmt.Println("bool:", *boolFlag)
	fmt.Println("string:", *stringFlag)

}
