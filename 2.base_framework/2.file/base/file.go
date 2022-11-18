package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Write 写文件
func Write(data []byte, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

// ReadFile 一次性全部读完
func ReadFile(filePath string) {

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", file)

}

// ReadFile2 一次性全部读完
func ReadFile2(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	fmt.Printf("%s", file)
}

// CopyFile 移动文件
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
