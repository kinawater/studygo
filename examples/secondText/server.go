package main

import "net/http"

type Server interface {
	RoutTable
	Start(address string) error
}
type sdkHttpServer struct {
	Name    string
	Handler Handler
	root    Filter
}

// Route 路由注册
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(wbh *WebHandle)) {
	s.Handler.Route(method, pattern, handlerFunc)
}

// Start 监听端口启动http服务器
func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c
		s.root()
	})
	http.Handle("/", s.Handler)
	return http.ListenAndServe(address, nil)
}

// NewHttpServer 构造函数
func NewHttpServer(name string, builders ...FilterBuilder) Server {
	// 声明一个路由map
	handler := NewRouteHandle()
	// 构造一个root根
	var root Filter = func(c *WebHandle) {
		handler.ServeHTTP(c.W, c.R)
	}
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		Handler: handler,
		root:    root,
	}
}
