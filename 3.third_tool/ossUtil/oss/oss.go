package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var bucketsClient *oss.Client

// GetBucketsClient 用来操作桶的
func GetBucketsClient() *oss.Client {
	return bucketsClient
}

func parse() (enpoint, key, secret string) {
	cfg, err := ini.Load("/Users/anthony/Documents/GitHub/study-go/3.third_tool/ossUtil/config.ini")
	if err != nil {
		log.Fatal("文件打开错误: ", err)
	}

	enpoint = cfg.Section("").Key("enpoint").String()
	key = cfg.Section("").Key("key").String()
	secret = cfg.Section("").Key("secret").String()
	return
}

func init() {
	fmt.Println("oss客户端初始化")
	enpoint, key, secret := parse()
	cli, err := oss.New(enpoint, key, secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketsClient = cli
}
