package main

import "fmt"



type  A struct{
    kuuyee int
}

func DoSomething(a *A){
	b :=a
	fmt.Printf("b address : %p\n",b)
}

func DoSomething1(a A){
	b := &a
	fmt.Printf("b1 address:%p",b)
}

func main() {
	var a *A
	var b A
	DoSomething(a)
	DoSomething1(b)

	
}
