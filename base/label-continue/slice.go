package main

import "fmt"

func updateSlice(s []int)  {
	s[0]=100
}
func main() {
	arr :=[...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]",arr[2:6])

	s1:=arr[2:6]
	fmt.Println(s1)
	s2:=s1[0:6]
	fmt.Println(s2)
}
