package main

import (
	"fmt"
	"sync"
	"time"
)

type Visited struct {
	mu sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count ++
	v.visited[url] = count
	return count
}
func visitedByUrl(urlList []string,visitedOpreat *Visited,c chan map[string]int) {
	for _,urlItem:=range urlList {
		c<-map [string] int{
			urlItem:visitedOpreat.VisitLink(urlItem),
		}
	}
}
func printC(c chan map[string]int){
	for item :=range c{
		fmt.Println(item)
	}
	close(c)
}
func main(){
	urlList := []string{
		"www.baidu.com",
		"www.baidu.com",
		"www.baidu.com",
		"www.souhu.com",
		"www.yahoo.com",
		"www.google.com",
		"www.toutiao.com",
	}
	visitedOpreat := Visited{visited: make(map [string] int)}
	c := make(chan map[string]int)
	go visitedByUrl(urlList,&visitedOpreat,c)

	go printC(c)
	time.Sleep(2*time.Second )
	//for item :=range c{
	//	fmt.Println(item)
	//}
	//for {
	//	select {
	//	case itme := <-c:
	//		fmt.Println(itme,"***")
	//	}
	//	case
	//}



}