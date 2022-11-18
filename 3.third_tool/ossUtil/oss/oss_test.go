package oss

import (
	"bufio"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"testing"
)

func TestGetCount(t *testing.T) {
	GetCount()
}

func TestListCount(t *testing.T) {
	//ListObjects()
	slice1 := []string{"1", "2", "3", "6", "8", "9", "10", "11", "a", "b", "c", "d", "e", "f", "g"}
	slice2 := []string{"2", "3", "5", "0"}

	in := intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)
	di := difference(slice1, slice2)
	fmt.Println("slice1与slice2的差集为：", di)

}

func TestListCount2(t *testing.T) {

	arr := result1()
	fmt.Println("bblog的长度是:", len(arr))

	arr = removeDuplicateElement(arr)
	fmt.Println("去重后的的长度是:", len(arr))

	strings := result2()
	tmp := len(strings)
	fmt.Println("阿里云的长度是:", len(strings))

	for _, v := range arr {
		if v != "" {
			for i := 0; i < len(strings); i++ {

				// 取值
				key := strings[i]

				if key == v && key != "" {
					strings = append(strings[:i], strings[i+1:]...)
					i--
				}

			}
		}

	}

	fmt.Println("剩余长度", len(strings))
	fmt.Println(tmp-len(arr) == len(strings))

	for _, k := range strings {
		DeleteObject(k)
	}

}

// 去重
func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}

			result = append(result, item)
		} else {
			fmt.Println(item)
		}
	}
	return result
}

func result1() (arr []string) {

	arr = make([]string, 0, 1000)
	filePath := "/Users/anthony/Documents/GitHub/study-go/3.third_tool/ossUtil/list.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR, 0666)
	defer file.Close()

	buf := bufio.NewReader(file)

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

		arr = append(arr, string(line))
	}
}

func result2() (arr []string) {
	arr = make([]string, 0, 1000)
	lists, err := GetObjectsClient().ListObjects(oss.MaxKeys(1000))
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range lists.Objects {
		arr = append(arr, v.Key)
	}
	return arr
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int, 1000)
	nn := make([]string, 1000)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int, 1000)
	nn := make([]string, 1000)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn

}
