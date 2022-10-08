package main

import (
	"fmt"
	"net"
)

func handlerTcp(conn net.Conn) {
	var body [100]byte
	defer conn.Close()
	for {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("当前信息:%s", string(body[:]))
		_, err = conn.Read(body[:])
		if err != nil {
			break
		}
	}

}

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handlerTcp(conn)
	}
}
