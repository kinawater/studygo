package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		fmt.Printf("client msg is:%s", msg)
		_, err = conn.Write([]byte("you send msg:" + msg))
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
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handlerConnection(conn)
	}
}
