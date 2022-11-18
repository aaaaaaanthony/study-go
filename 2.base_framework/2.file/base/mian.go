package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
