package main

import "fmt"

func main() {

	s := "Hello沙河"

	// len()求得是byte字节的数量
	n := len(s) // q求字符串s的长度,吧长度保存到变量n中
	fmt.Println(n)

	for k, v := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%v ,%v , %c\n", k, v, v)
	}

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把runne切片强制转换成字符串

	// TODO 看不太懂这里
	c1 := "红"
	c2 := '红' // rune(int32)
	fmt.Printf("c1:%T c2:%T \n", c1, c2)

	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c1:%T c2:%T \n", c3, c4)
	fmt.Printf("%d\n", c4)

	// 类型转换
	n1 := 10 //int
	var f float64
	f = float64(n1)
	fmt.Printf("n1:%T f:%T \n", n1, f)
}
