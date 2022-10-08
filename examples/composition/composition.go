package main

import "fmt"

// 游泳技能
type Swimming interface {
	Swim()
}

//鸭子
type Duck interface {
	// 组合游泳
	Swimming
}

// 组合2
type Concrete2 struct {
	*Base
}

type Base struct {
	Name string
}

// 组合实现1
type Concrete1 struct {
	Base
}

func (c *Concrete1) SayHello() {
	// 打印 c.Name
	fmt.Println("name is : #{c.Name}")
	fmt.Println("222name is : #{c.Name} ")
	// 调用被组合的
	c.Base.SayHello()
	//调用自己
	c.SayHello()
	c.Base.SayGoodBye()
}

// base结构的sayhello方法
func (b *Base) SayHello() {
	fmt.Println("base 的 sayhello，name is #{b.Name}")
}

func (b *Base) SayGoodBye() {

}
