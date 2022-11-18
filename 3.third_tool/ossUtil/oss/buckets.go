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

// GetCount 获取文件数量
func GetCount() {

	stat, _ := GetBucketsClient().GetBucketStat("hexosrc")

	// 获取Bucket的总存储量，单位为字节。
	fmt.Println("Bucket Stat Storage:", stat.Storage)
	// 获取Bucket中总的Object数量。
	fmt.Println("Bucket Stat Object Count:", stat.ObjectCount)
	// 获取Bucket中已经初始化但还未完成（Complete）或者还未中止（Abort）的Multipart Upload数量。
	fmt.Println("Bucket Stat Multipart Upload Count:", stat.MultipartUploadCount)
	// 获取Bucket中Live Channel的数量。
	fmt.Println("Bucket Stat Live Channel Count:", stat.LiveChannelCount)
	// 此次调用获取到的存储信息的时间点，格式为时间戳，单位为秒。
	fmt.Println("Bucket Stat Last Modified Time:", stat.LastModifiedTime)
	// 获取标准存储类型Object的存储量，单位为字节。
	fmt.Println("Bucket Stat Standard Storage:", stat.StandardStorage)
	// 获取标准存储类型的Object的数量。
	fmt.Println("Bucket Stat Standard Object Count:", stat.StandardObjectCount)
	// 获取低频存储类型Object的计费存储量，单位为字节。
	fmt.Println("Bucket Stat Infrequent Access Storage:", stat.InfrequentAccessStorage)
	// 获取低频存储类型Object的实际存储量，单位为字节。
	fmt.Println("Bucket Stat Infrequent Access Real Storage:", stat.InfrequentAccessRealStorage)
	// 获取低频存储类型的Object数量。
	fmt.Println("Bucket Stat Infrequent Access Object Count:", stat.InfrequentAccessObjectCount)
	// 获取归档存储类型Object的计费存储量，单位为字节。
	fmt.Println("Bucket Stat Archive Storage:", stat.ArchiveStorage)
	// 获取归档存储类型Object的实际存储量，单位为字节。
	//fmt.Println("Bucket Stat Archive Real Storage:", stat.ArchiveRealStorage)
	// 获取归档存储类型的Object数量。
	fmt.Println("Bucket Stat Archive Object Count:", stat.ArchiveObjectCount)
	// 获取冷归档存储类型Object的计费存储量，单位为字节。
	fmt.Println("Bucket Stat Cold Archive Storage:", stat.ColdArchiveStorage)
	// 获取冷归档存储类型Object的实际存储量，单位为字节。
	fmt.Println("Bucket Stat Cold Archive Real Storage:", stat.ColdArchiveRealStorage)
	// 获取冷归档存储类型的Object数量。
	fmt.Println("Bucket Stat Cold Archive Object Count:", stat.ColdArchiveObjectCount)

}
