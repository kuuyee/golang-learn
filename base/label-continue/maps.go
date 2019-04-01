package main

import "fmt"

func main() {
	m:=map[string]string{
		"kuuyee":"光子",
		"jiangxuguang":"姜旭光",
	}

	fmt.Println(m)
	for k,v :=range m{
		fmt.Println(k,v)
	}


	if kuuyee,ok :=m["kuuyee"];ok{
		fmt.Println(kuuyee)
	}else {
		fmt.Println("不存在")
	}

}
