package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p ="ciao"

	fmt.Printf("Here is the pointer p: %p\n",p)
	fmt.Printf("Here is the string *p %s\n",*p)
	fmt.Printf("Here is the string s: %s",s)

	const i = 5
	//ptr := &i  //error, 不能获取常量的地址
}
