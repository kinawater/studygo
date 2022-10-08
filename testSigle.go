package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP:
				fmt.Println("Program Exit...", s)
				fmt.Println("SIGHUP")
				GracefullExit()
			case syscall.SIGINT:
				fmt.Println("Program Exit...", s)
				fmt.Println("SIGHUP")
				GracefullExit()
			case syscall.SIGTERM:
				fmt.Println("Program Exit...", s)
				fmt.Println("SIGHUP")
				GracefullExit()
			case syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				fmt.Println("SIGHUP")
				GracefullExit()
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("Program Start...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func GracefullExit() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}
