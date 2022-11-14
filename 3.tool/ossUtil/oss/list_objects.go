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

// GetObjectsClient 用来操作桶里的文件的
func GetObjectsClient() (bucket *oss.Bucket) {
	bucket, err := GetBucketsClient().Bucket("hexosrc")
	if err != nil {
		fmt.Println(err)
	}
	return
}
