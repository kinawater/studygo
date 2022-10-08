package main

import "fmt"

func main() {
	son := Son{
		Parent{},
	}
	son.SayName()
	son.Parent.SayName()
}

type Parent struct {
}

func (p *Parent) SayName() {
	fmt.Println(p.Name())
}
func (p *Parent) Name() string {
	return "parent"
}

type Son struct {
	Parent
}

// son重写Name方法
func (s *Son) Name() string {
	return "son"
}
