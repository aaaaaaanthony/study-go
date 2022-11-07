package mynet

import (
	"fmt"
	"net"
)

func SplitHostPort(addr string) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		fmt.Printf("报错了:%s\n", err)
	}
	fmt.Printf("IP:%s,端口:%s", host, port)
}
