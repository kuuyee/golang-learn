package main

import (
	"fmt"
)

func main() {
	//b := [512]byte{'x', 'y', 'z'} //数组不能string([]byte)转换操作
	b := []byte{'x', 'y', 'z'} //切片是可以string([]byte)转换
	fmt.Println(string(b))
}
