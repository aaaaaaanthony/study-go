package main

import (
	"study-go/3.tool/ossUtil/oss"
)

func main() {

	// 桶列表
	oss.ListBuckets()

	// 文件列表,不知道为啥没有显示完整
	oss.ListObjects()

	// 删除文件
	//key := "202211141302816.png"
	//oss.DeleteObject(key)

}
