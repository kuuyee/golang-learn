package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("错误的地址")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("地址：", addr.String())
	fmt.Println("默认掩码长度：", bits)
	fmt.Println("Leading ones count is：", ones)
	fmt.Println("掩码(hex)：", mask.String())
	fmt.Println("网络：", network.String())
	os.Exit(0)
}
