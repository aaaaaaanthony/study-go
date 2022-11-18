package main

import (
	"fmt"
	"net"
	"sync"
)

// ResolveIPAddr
// ResolveIPAddr将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析。
// 函数会在参数net指定的网络类型上解析，net必须是"ip"、"ip4"或"ip6"。
func ResolveIPAddr() {

	addr := "ip"
	ipAddr, err := net.ResolveIPAddr(addr, "baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ipAddr)
}

// SplitHostPort 解析端口和IP
func SplitHostPort() {

	addr := "127.0.0.1:8080"

	host, p, err := net.SplitHostPort(addr)
	fmt.Println(host, p, err)

}

// Dial 在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
// "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
// 对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。
var wg sync.WaitGroup

func Dial() {

	wg.Add(1)
	go func() {
		listen, err2 := net.Listen("tcp", "127.0.0.1:8089")
		if err2 != nil {
			fmt.Println(err2)
		}

		a := listen.Addr()
		fmt.Println("服务端信息Network:", a.Network())
		fmt.Println("服务端信息String:", a.String())

		for {
			accept, _ := listen.Accept()

			fmt.Println("服务端接收信息LocalAddr,Network:", accept.LocalAddr().Network())
			fmt.Println("服务端接收信息LocalAddr,Network:", accept.LocalAddr().String())
			fmt.Println("服务端接收信息RemoteAddr,Network:", accept.RemoteAddr().Network())
			fmt.Println("服务端接收信息RemoteAddr,String:", accept.RemoteAddr().String())

			// 接收消息
			buf := make([]byte, 1024)
			n, err := accept.Read(buf[0:])
			if err != nil {
				return
			}
			fmt.Printf("服务端接收的消息 : %v\n", string(buf[:n]))

			// 发送消息
			sendMsg := "hello client"
			n, err = accept.Write([]byte(sendMsg))
			if err != nil {
				panic("客户端发送消息失败")
			} else {
				fmt.Printf("客户端发送消息成功,内容是:%s,长度是:%d \n", sendMsg, n)
			}

			wg.Done()
		}
	}()

	dial, err := net.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println(err)
		panic(err)

	}

	addr := dial.LocalAddr()
	fmt.Println("客户端信息LocalAddr.String:", addr.String())
	fmt.Println("客户端信息LocalAddr.Network:", addr.Network())

	remoteAddr := dial.RemoteAddr()
	fmt.Println("客户端信息RemoteAddr.Network:", remoteAddr.String())
	fmt.Println("客户端信息RemoteAddr.Network:", remoteAddr.Network())

	// 发送消息
	sendMsg := "hello server"
	n, err := dial.Write([]byte(sendMsg))
	if err != nil {
		panic("客户端发送消息失败")
	} else {
		fmt.Printf("客户端发送消息成功,内容是:%s,长度是:%d \n", sendMsg, n)
	}

	buf := make([]byte, 1024)
	n, err = dial.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Printf("服务端接收的消息 : %v\n", string(buf[:n]))
	wg.Wait()
}
