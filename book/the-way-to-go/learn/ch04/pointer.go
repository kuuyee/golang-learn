package main

import (
	"fmt"
)

func main() {

	var i1 =5
	fmt.Printf("一个整数： %d, 其内存地址是： %p\n",i1,&i1)
	var i2 =5
	fmt.Printf("另一个整数： %d, 其内存地址是： %p\n",i2,&i2)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n",intP,*intP)
}
