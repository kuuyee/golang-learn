package main

import "fmt"

func main() {
	F1("3, 4, 5", "6", "67")
}

func F1(s ...string) {
	F2(s...)
}

func F2(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
