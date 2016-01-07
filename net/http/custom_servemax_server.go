package main

import (
	"io"
	"log"
	"net/http"
)

// 定义一个map存储url及对应的处理函数
var myServeMax map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 定义路由匹配逻辑
	if h, ok := myServeMax[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func main() {
	// 第一个http.Server结构
	server := http.Server{
		Addr:    ":7777",
		Handler: &myHandler{}, // 使用自定的ServeHTTP方法，进入自定义路由逻辑
	}

	myServeMax = make(map[string]func(w http.ResponseWriter, r *http.Request))
	myServeMax["/foo"] = fooFunc
	myServeMax["/bar"] = barFunc

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func fooFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "路由执行函数：fooFunc")
}
func barFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "路由执行函数：barFunc")
}
