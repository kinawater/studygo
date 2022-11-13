package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"

	fmt.Println(m["1"])
}
