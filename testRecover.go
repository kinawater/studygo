package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start...")
	var n = 0
	for true {
		levelOne(n)
		n++
	}
	fmt.Println("...END")

}

func levelOne(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("in the recover")
			fmt.Println(err)
			fmt.Println("system is recover")
		}
	}()

	for true {
		fmt.Println("now n is " + strconv.Itoa(n))
		time.Sleep(time.Second * 1)
		panic("haha down")

	}
}
