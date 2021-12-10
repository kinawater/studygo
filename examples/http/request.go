package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
func main(){
	http.HandleFunc("/",home)
	http.HandleFunc("/body/once",readBodyOnce)
	http.HandleFunc("/body/repeat",getBodyReadRepeat)
	http.HandleFunc("/body/query",queryParams)
	http.HandleFunc("/body/form",form)
	http.HandleFunc("/body/header",header)
	http.HandleFunc("/body/wholeUrl",wholeUrl)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
