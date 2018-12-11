package main

import "fmt"

func main() {
	const (
		B  float64 = 1 << (iota * 10)
		KB         = 1 << (iota * 10)
		MB         = 1 << (iota * 10)
		GB         = 1 << (iota * 10)
		TB         = 1 << (iota * 10)
		PB         = 1 << (iota * 10)
		EB         = 1 << (iota * 10)
		ZB         = 1 << (iota * 10) //constant 1208925819614629174706176 overflows int
		YB         = 1 << (iota * 10) //constant 1208925819614629174706176 overflows int
	)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	//fmt.Println(ZB)
	//fmt.Println(YB)

}
