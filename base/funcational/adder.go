package main

import "fmt"

func adder() func(int) int  {

}

func main() {
	a:=adder()
	for i:=0;i<10 ;i++  {
		fmt.Printf(a(i))
	}
}
