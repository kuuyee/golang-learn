package main

import (
	"fmt"
	"math"
)

var Pi float64
func init(){
	Pi = 4*math.Atan(1)
	fmt.Printf("Pi的值是：%f\n", Pi)
}

func other() {
	fmt.Printf("不调用不会执行")
}

func main() {
	fmt.Printf("在init函数之后执行！")
}
