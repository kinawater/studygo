package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContentText struct {
	W http.ResponseWriter
	R *http.Request
}
func (c *ContentText) ReadJson(obj interface{}) error{
	//读取body，反序列化
	body,err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body,obj)
	if err != nil {
		return err

	}
	return nil
}
func (c *ContentText) WriteJson(code int,obj interface{}) error {
	// 序列化输出息结果
	c.W.WriteHeader(code)
	respJson,err := json.Marshal(obj)
	if err != nil {
		fmt.Fprintf(c.W,"序列化数据失败 %v",err)
		return err
	}
	_,err = c.W.Write(respJson)
	return err
}
func (c *ContentText) BadRequest(obj interface{}) error {
	return c.WriteJson(http.StatusBadRequest,obj)
}
func NewContext(w http.ResponseWriter,r *http.Request) *ContentText{
	return &ContentText{
		W: w,
		R: r,
	}
}