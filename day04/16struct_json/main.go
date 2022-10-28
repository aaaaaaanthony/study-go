package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age`
}

// 把结构体转成json
func main() {

	p1 := person{
		Name: "anthony",
		Age:  1,
	}

	// 序列化
	value, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value))

	// 反序列化
	str := `{"name":"anthony","age":13}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)

}
