package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("解析错误", err.Error())
		os.Exit(1)
	}

	fmt.Println("解析到的地址为：", addr.String())
	os.Exit(0)
}
