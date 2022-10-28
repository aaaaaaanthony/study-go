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

}
