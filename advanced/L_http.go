package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
在这个例子中，类型 Hello 实现了 http.Handler。

访问 http://localhost:4000/ 会看到来自程序的问候。
*/

type WebServer struct{}

func (web WebServer) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Http!</h1>")
}

func main() {
	var s WebServer
	err := http.ListenAndServe("127.0.0.1:10086", s)
	if err != nil {
		log.Fatal(err)
	}
}
