package main

import my "study-go/mydoc/mynet"

func main() {

	addr := "127.0.0.1:8110"
	my.SplitHostPort(addr)
}
