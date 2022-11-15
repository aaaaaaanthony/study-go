package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"study-go/3.third_tool/ossUtil/oss"
)

func main() {

	for {
		var key string
		fmt.Println("=====输入文件名===========")
		fmt.Scan(&key)

		key = strings.Split(key, ")")[0]
		fmt.Println(key)
		myrange(key)
		fmt.Println("一轮完成")
	}

}

func myrange(key string) {
	// 桶列表
	//oss.ListBuckets()

	// 文件列表,不知道为啥没有显示完整
	//oss.ListObjects()

	// 删除文件

	var number int = funcName() + 1
	fileName := strconv.Itoa(number) + ".png"
	//oss.DeleteObject(key)

	// 文件是否存在
	exist := oss.IsObjectExist(key)
	fmt.Println(exist)
	if exist {

		// 下载文件

		oss.Download(key, fileName)

		// 删除文件
		oss.DeleteObject(key)

		src := "/Users/anthony/Documents/GitHub/study-go/" + fileName
		des := "/Users/anthony/Documents/GitHub/notes/assets/" + fileName
		written, err := CopyFile(src, des)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(written)
		}
	}

	fmt.Println(funcName())
}

func funcName() (max int) {
	path := "/Users/anthony/Documents/GitHub/notes/assets/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	max = 0

	// 读取 dirname 指定的目录， 并返回一个根据文件名进行排序的目录节点列表。
	for _, file := range files {

		split := strings.Split(file.Name(), ".")
		numberString := split[0]
		atoi, _ := strconv.Atoi(numberString)

		if atoi > max {
			max = atoi
		}
	}

	return
}

func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
