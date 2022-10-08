package main

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

//添加
func (c *TaskController) Add() {
	c.Ctx.WriteString("add")
}

// 查询
func (c *TaskController) Query() {
	c.Ctx.WriteString("Query")
}

// 删除
func (c *TaskController) Del() {
	c.Ctx.WriteString("Del")
}

// 修改
func (c *TaskController) Modify() {
	c.Ctx.WriteString("Modify")
}

func main() {
	beego.Router("/task/", &TaskController{}, "get:Add;get,post:Query")
	beego.Run()
}
