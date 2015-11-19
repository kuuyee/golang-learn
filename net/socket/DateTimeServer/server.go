package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("时间服务器启动。。。监听7777端口")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("收到1个请求！addr:", conn.RemoteAddr())

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "致命错误：", err.Error())
	}
}
