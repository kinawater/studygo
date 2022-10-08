package main

import "fmt"

func main() {

	s := []byte("abcdefg\r\n")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("@@@" + string(s[len(s)-1]) + "@@@")
}
