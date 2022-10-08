package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// 构造处理类型
type WebHandle struct {
	W http.ResponseWriter
	R *http.Request
}

// 创建一个新webHandle对象
func NewWebHandle(w http.ResponseWriter, r *http.Request) *WebHandle {
	return &WebHandle{
		W: w,
		R: r,
	}
}

// 读取json
func (wh *WebHandle) ReadJson(obj interface{}) error {
	// 读取obj，对obj进行反序列化
	body, err := io.ReadAll(wh.R.Body)
	if err != nil {
		return err
	}
	// 反序列化,把body里面的参数塞入到obj里面
	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return err
}

// 写入并以json的形式发送
func (wh *WebHandle) WriteJson(code int, obj interface{}) error {
	// 写入状态码code
	wh.W.WriteHeader(map[bool]int{IsNil(code): code, !IsNil(code): http.StatusNotFound}[false])
	// 把对象转换为json
	respJson, err := json.Marshal(obj)
	if err != nil {
		fmt.Fprintf(wh.W, "json转换失败 %v", err)
		return err
	}
	_, err = wh.W.Write(respJson)
	return nil
}

// 错误路由
func (wh *WebHandle) badRoute() {
	wh.WriteJson(http.StatusNotFound, map[string]string{"status": "404", "msg": "路由不存在"})
}

// 反射判断对象是否为空
func IsNil(obj interface{}) bool {
	valueInterface := reflect.ValueOf(obj)
	if valueInterface.Kind() == reflect.Ptr {
		return valueInterface.IsNil()
	}
	return false
}
