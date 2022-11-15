package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {

	str := "![](https://image.mmzcg.com/blog/copy_20210114203327.png)"
	// 配置md的图片格式
	matched, _ := regexp.MatchString("!\\[\\]\\(.*\\)", str)
	if matched {
		// 只是匹配地址
		compile, _ := regexp.Compile("https://.*.png")
		url := compile.FindString(str)
		fmt.Println(url)

		// 图片下载下来

		// 找打图片路径
		split := strings.Split(url, "image.mmzcg.com")
		fmt.Println(split[1])
	}

	//funcName()
}

// 下载图片信息
func downLoad(base string, url string) error {
	pic := base
	idx := strings.LastIndex(url, "/")
	if idx < 0 {
		pic += "/" + url
	} else {
		pic += url[idx+1:]
	}
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		return err
	}
	err = ioutil.WriteFile(pic, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return err
	}
	return nil
}

func funcName() {
	path := "/Users/anthony/Documents/GitHub/study-go/2.base_framework/2.file/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	// 读取 dirname 指定的目录， 并返回一个根据文件名进行排序的目录节点列表。
	for _, file := range files {
		fmt.Println(file.Name())

		if file.Name() == "test.txt" {
			output, flag, err := line(path + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// 重写文件
			if flag == true {
				err := write(output, path+file.Name())
				if err != nil {
					fmt.Println("写入文件失败", err)
				} else {
					fmt.Println("写入文件成功")
				}
			}
		}
	}
}

func write(data []byte, filePath string) error {
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

// 一次性全部读完
func readFile(filePath string) {

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", file)

}

// 一次性全部读完
func readFile2(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	fmt.Printf("%s", file)
}

func line(filePath string) (output []byte, replaceFlag bool, err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	// 输出完整内容
	output = make([]byte, 0)

	// 是否要替换文件
	replaceFlag = false

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				return output, replaceFlag, nil
			} else {
				fmt.Println("Read file error!", err)
				return output, replaceFlag, err
			}
		}

		fmt.Println("行内容===>", line)

		// 匹配       !\[\]\(.*\)
		// ![](https://image.mmzcg.com/blog/copy_20210114203327.png)
		matched, _ := regexp.MatchString("Anthony", string(line))

		if matched {
			replaceString := strings.Replace(string(line), "Anthony", "Hello", -1)
			output = append(output, []byte(replaceString)...)
			output = append(output, []byte("\n")...)

			// 标识要替换
			if replaceFlag == false {
				replaceFlag = true
			}

		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}

	}

	return output, replaceFlag, nil

}
