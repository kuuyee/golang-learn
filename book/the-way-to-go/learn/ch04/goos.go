package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS;
	fmt.Printf("当前操作系统是：%s\n",goos)

	path :=os.Getenv("PATH")
	fmt.Printf("当前PAHT变量值是 %s\n",path)
}
