package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func main() {

	go hello()
	fmt.Println("main")
	time.Sleep(time.Second)
}
