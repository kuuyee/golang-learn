package main

import "fmt"

type number struct {
	f float32
}

type nr number //别名类型

func main() {
	a := number{5.0}
	b := nr{5.0}
	var c = number(b)
	fmt.Println(a, b, c)
}
