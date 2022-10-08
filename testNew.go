package main

import (
	"bytes"
	"fmt"
)

var CRLF = "\r\n"

func main() {
	a := StatusReply{Status: "1"}
	b := new(StatusReply)
	b.Status = "1"
	c := *b
	c.Status = "2"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(c)
}

type StatusReply struct {
	Status string
}

func (r *StatusReply) ToBytes() []byte {
	var buf bytes.Buffer
	buf.WriteString("+" + r.Status + CRLF)
	return buf.Bytes()
}
