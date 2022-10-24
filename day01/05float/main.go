package main

import (
	"fmt"
)

// 浮点数
func main() {
	// math.MaxFloat32   float32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // go默认小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)

	// f1= f2  float32类型的值不能直接转
}
