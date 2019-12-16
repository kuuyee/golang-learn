package main

import (
	"fmt"
)

func main() {
	input(2)
}

func input(num int) (old1 int, old2 int) {

	old1 = num - 1
	old2 = num - 2
	return computed(num, old1, old2)
}

func computed(num int, old1 int, old2 int) (result int, next int) {
	result = old1 + old2
	fmt.Printf("当前输入为 %d , 计算结果为 %d\n", num, result)
	result++
	return input(result)
}
