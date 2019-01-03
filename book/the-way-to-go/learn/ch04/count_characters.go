package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("这个字符串的字节数量是 %d\n",len(s1))
	fmt.Printf("这个字符串的字符数量是 %d\n",utf8.RuneCountInString(s1))

	s2 := "asSASA ddd dsjkdsjsăă dk"
	fmt.Printf("这个字符串的字节数量是 %d\n",len(s2))
	fmt.Printf("这个字符串的字符数量是 %d\n",utf8.RuneCountInString(s2))
}
