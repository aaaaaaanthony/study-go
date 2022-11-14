package main

import (
	"fmt"
	"net/http"
	my "study-go/1.study-base/mydoc/mynet"
)

func main() {

	//addr := "127.0.0.1:8110"
	//my.SplitHostPort(addr)

	//http.HandleFunc("/path", f2)

	my.HandleFunc("/path", f1)

}

func f1(request string, response int) {
	fmt.Println("我自己的名字时:", request, "年纪是:", response)

}

func f2(http.ResponseWriter, *http.Request) {
	fmt.Println("我自己的")

}
