package main

import (
	"fmt"
)

type foo struct {
	name string
}

func NewFoo() *foo {
	fmt.Println("立即被执行了！")
	return &foo{name: "KuuYee"}
}

func (f *foo) HandleF() {
	fmt.Println("处理！：", f.name)
}

// NewFoo()在Main直行之前就被直行了
var DefaultFoo = NewFoo()

func main() {
	fmt.Println("开始直行Main()")
	DefaultFoo.HandleF()
}
