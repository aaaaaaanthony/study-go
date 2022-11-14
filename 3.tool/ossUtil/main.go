package main

import (
	"fmt"
	"study-go/3.tool/ossUtil/oss"
)

func main() {

	// 桶列表
	//oss.ListBuckets()

	// 文件列表,不知道为啥没有显示完整
	//oss.ListObjects()

	// 删除文件
	//key := "16547418.jpg"
	key := "blog/copy_20201226164302.png"
	//oss.DeleteObject(key)

	// 文件是否存在
	exist := oss.IsObjectExist(key)
	fmt.Println(exist)

}
