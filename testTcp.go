package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	ac, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	// 设置一个接受缓存
	var body [100]byte

	for true {
		_, err := ac.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到消息:%s", string(body[:]))
		_, err = ac.Write(body[:])
		if err != nil {
			break
		}
	}
}
