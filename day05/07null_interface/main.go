package main

import "fmt"

// 空接口
func main() {

	m1 := make(map[string]interface{})
	m1["k1"] = "v1"
	m1["k2"] = 123

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	assign2("123")

}

// 类型断言1
func assign(x interface{}) {

	fmt.Printf("%T\n", x)
	str, ok := x.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("才对了", str)
	}

}

// 类型断言2
func assign2(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Println("string", v)
	}

}
