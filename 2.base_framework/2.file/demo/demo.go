package demo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var path2 = "/Users/anthony/Documents/GitHub/blog/source/"

func funcName() {

	fileCry(path2)
}

func fileCry(path string) {
	files, _ := ioutil.ReadDir(path)
	// 读取 dirname 指定的目录， 并返回一个根据文件名进行排序的目录节点列表。
	for _, file := range files {

		// 文件名或者目录名
		fileName := file.Name()

		if file.IsDir() {
			// 目录
			fileCry(path + "/" + fileName)
		} else {
			// 文件
			if matched, _ := regexp.MatchString(".*.md", fileName); matched == true {
				line(path + "/" + fileName)
			}
		}

	}
}

func line(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_RDWR, 0666)
	defer file.Close()

	buf := bufio.NewReader(file)

	count := 0

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				//fmt.Println("File read ok!")
				return
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}

		// 匹配
		if matched, _ := regexp.MatchString("image.mmzcg.com/.*/.*", string(line)); matched {
			fmt.Println(string(line))
			count++
		}
	}
	fmt.Println(count)
}
