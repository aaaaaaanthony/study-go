package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// ListObjects 文件列表
func ListObjects() {
	lists, err := GetObjectsClient().ListObjects()
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range lists.Objects {
		fmt.Printf("序号:%v,文件名:%v\n", k, v.Key)
	}

}

// DeleteObject 删除文件
func DeleteObject(key string) {

	err := GetObjectsClient().DeleteObject(key)
	if err != nil {
		fmt.Println("删除文件失败")
	}
	fmt.Println("删除文件成功")

}

// IsObjectExist 判断文件是否存在
func IsObjectExist(key string) (exist bool) {

	exist, err2 := GetObjectsClient().IsObjectExist(key)
	if err2 != nil {
		fmt.Println("查找文件错误")
	}

	if exist == true {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}

	return
}

// 下载文件
func Download(key string, newName string) {
	err := GetObjectsClient().GetObjectToFile(key, newName)
	if err != nil {
		fmt.Println("文件下载失败")
	} else {
		fmt.Println("文件下载成功")
	}

}

// GetObjectsClient 用来操作桶里的文件的
func GetObjectsClient() (bucket *oss.Bucket) {
	bucket, err := GetBucketsClient().Bucket("hexosrc")
	if err != nil {
		fmt.Println(err)
	}
	return
}
