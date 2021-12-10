package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type signUpReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type commonResponse struct {
	BizCode int `json:"biz_code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
func home(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"hello,minasa,I love %s!",r.URL.Path[1:])
}
// 只读一次body测试
func readBodyOnce(w http.ResponseWriter,r *http.Request){
	body,err:=io.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"read body failed: %v",err)
		return
	}
	fmt.Fprintf(w,"read body have data:%s \n",string(body))
	//第二次读
	body2, err2 := io.ReadAll(r.Body)
	if err !=nil {
		fmt.Fprintf(w,"read the data one more time",err2)
	}
	fmt.Fprintf(w,"read the data one more time:[%s],and data length %d \n",string(body2),len(body2))
}
// 重复读body  getbody
func getBodyReadRepeat(w http.ResponseWriter,r *http.Request){
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}

}
//获取url链接上的参数
func queryParams(w http.ResponseWriter,r * http.Request){
	values := r.URL.Query()
	fmt.Fprintf(w,"query is %v \n",values)
}
func form(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"没有进行 parse form之前的数据 %v\n",r.Form)
	err := r.ParseForm()
	if err !=nil {
		fmt.Fprintf(w,"parse解析处理后的form数据 %v \n",r.Form)
	}
	fmt.Fprintf(w,"没有进行 parse form之前的数据 %v\n",r.Form)
}
// header
func header(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"header is %v\n",r.Header)
}
func wholeUrl(w http.ResponseWriter,r *http.Request){
	data,_:= json.Marshal(r.URL)
	fmt.Fprintf(w,string(data))
}
func SignUpWithoutContext(w http.ResponseWriter,r *http.Request){
	req := &signUpReq{}
	ctx := &ContentText{
		W:w,
		R:r,
	}
	err := ctx.ReadJson(req)
	if err !=nil {
		fmt.Fprintf(w,"反序列化失败deserialized failed: %v",err)
		return
	}
	resp := &commonResponse{
		Data:123,
	}
	err = ctx.WriteJson(http.StatusOK,resp)
	if err!=nil{
		fmt.Printf("写入失败")
	}
}
func SignUpHaveContext(ctx *ContentText){
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err !=nil {
		ctx.BadRequest(err)
		return
	}
	resp := &commonResponse{
		Data:123,
	}
	err = ctx.WriteJson(http.StatusOK,resp)
	if err!=nil{
		fmt.Printf("写入失败")
	}
}


func main(){
	server := NewHttpServer("testServer")
	server.Route("/user/signUp",SignUpHaveContext)
	if err := server.Start(":8080"); err != nil {
		panic(err)
	}
}
