package main

import (
	"fmt"
	"net/http"
)

type aaaa struct {
	client *http.Client
}

func main() {
	var b [][]byte
	b = make([][]byte, 0)
	b = append(b, []byte("abc"))
	b = append(b, []byte("哈哈哈哈"))
	for _, bytes := range b {
		fmt.Println(string(bytes))
	}
	//fmt.Println(b)
}
