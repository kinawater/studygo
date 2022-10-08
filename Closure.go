package main

import "fmt"

func main() {
	str := "how are you"
	fmt.Println(str)
	foo := func() {
		// 匿名函数访问外部变量
		str = "fine"
		fmt.Println(str)
	}
	fmt.Println("3", str)
	foo()
}
