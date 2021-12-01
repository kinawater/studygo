package main

import (
	"fmt"
	"strings"
)

func main()  {
	upc := make(chan string)
	dwc := make(chan string)
	go sourceGopher(upc)
	go filterGopher(upc,dwc)
	printGopher(dwc)
}
func sourceGopher(downstream chan string) {
	for _, v:= range [] string{"hello oked","hi nope","moximoxi oked"}{
		downstream <- v
	}
	close(downstream)
}
/**
过滤不包含oked，并塞入新通道
 */
func filterGopher(upstream,downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item,"oked") {
			downstream <- item
		}
	}
	close(downstream)
}
func printGopher(upstream chan string) {
	for item:=range upstream{
		fmt.Println(item)
	}
}