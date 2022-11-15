package oss

import "fmt"

// ListBuckets 获取桶列表
func ListBuckets() {

	client := GetBucketsClient()
	lists, err := client.ListBuckets()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("桶列表")
	for _, buck := range lists.Buckets {
		fmt.Println("桶名", buck.Name)
	}
}
