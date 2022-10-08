package main

import (
	"fmt"
	"sync"
)

type DataEntity struct {
	Data interface{}
}

func main() {
	var m sync.Map
	testA := DataEntity{Data: 123}
	m.Store("a", testA)

	aValue, ok := m.Load("a")
	fmt.Println(aValue)
	fmt.Println(ok)
	testB := DataEntity{Data: 456}
	m.Store("a", testB)
	aValue, ok = m.Load("a")
	fmt.Println(aValue)
	fmt.Println(ok)
	//fmt.Printf("&mp:%p\n",&mp)
	//
	//
	//
	//
	//i := 0
	//for i<=10 {
	//	mp[i] = i
	//	fmt.Printf("mp1:%p mp2:%p\n",&mp,mp)
	//	fmt.Println(mp)
	//	i++
	//}
}
