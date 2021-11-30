package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i:=0;i<5;i++{
		go sleepyGopher2(i,c)
	}
	for i:=0;i<5;i++{
		gopherID := <-c
		fmt.Println("Now goperID",gopherID)
	}
}
func sleepyGopher2(i int,c chan int) {
	times := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(times)
	fmt.Println("哈哈哈哈哈",i,"-----时间",times)
	c <- i
}