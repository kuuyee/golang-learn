package main

import (
	"fmt"
	"strconv"
)

func main() {

	var orig string ="777"
	var an int
	var newS string

	fmt.Printf("The Size of ints is: %d\n",strconv.IntSize)

	an,_ = strconv.Atoi(orig)
	fmt.Printf("The integer si: %d\n",an)
	an = an+5

	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n",newS)
}
