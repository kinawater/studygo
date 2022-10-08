package main

import (
	"fmt"
	"net/http"
	"time"
)

// 注册信息
type signUpReq struct {
	MyEmail           string `json:"email"`
	PassWord          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

//输出信息
type outputResponse struct {
	Name    string      `json:"name"`
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 处理
func handelMsg(webHandle *WebHandle) {
	//解析 json
	nowSignUpReq := &signUpReq{}
	err := webHandle.ReadJson(nowSignUpReq)
	if err != nil {
		fmt.Fprintf(webHandle.W, "读取错误,%v", err)
		return
	}
	// 返回值
	resp := &outputResponse{
		Name:    nowSignUpReq.MyEmail,
		BizCode: 123,
		Msg:     "注册成功",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err = webHandle.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Fprintf(webHandle.W, "注册失败")
		return
	}
}
func main() {
	httpServer := NewHttpServer("第一个")
	httpServer.Route(http.MethodPost, "/user/signUp", handelMsg)
	if err := httpServer.Start(":8080"); err != nil {
		panic(err)
	}
}
