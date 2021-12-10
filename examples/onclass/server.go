package main

import "net/http"

type Server interface {
	Route(pattern string , handleFunc func(ctx *ContentText))
	Start(address string) error
}
type sdkHttpServer struct {
	
	Name string
}
// 路由注册
func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *ContentText)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer,request)
		handleFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address,nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}








//func NewServer() Server {
//	return &sdkHttpServer{}
//}
//type Factory func() Server
//
//var factory Factory
//func RegisterFactory(f Factory){
//	factory = f
//}
//func NewServer() Server{
//	return factory()
//}
