package main

import "fmt"

func watch(c chan int) {
	if <-c == 1 {
		fmt.Println("hello")
	}
}

func main() {
	c := make(chan int)
	go watch(c)
	c <- 1
}
