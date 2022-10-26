package main

import "fmt"

func main() {

	// 要估算该map容量,避免在程序运行期间再动态扩容
	var m1 = make(map[string]int, 10)

	m1["k1"] = 1

	fmt.Printf("长度:%d,容量:%d\n", len(m1), 1)
	fmt.Println(m1["k1"])

	// 判断有没有key
	value, ok := m1["k1"]
	if ok {
		fmt.Println("找到数据了", value)
	} else {
		fmt.Println("没有找到数据了", value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 删除
	delete(m1, "m2")
	fmt.Println(m1)
}
