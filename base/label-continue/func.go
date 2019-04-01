package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {

	return a / b, a % b
}

func sum(values ...int) int  {
	s:=0
	for i:=range values{
		s += values[i]
	}
	return s
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funcation %s with args "+"(%d, %d)", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func swap(a,b int)  {
	b,a = a,b
}

func main() {
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))


	fmt.Println(sum(1,2,3,4,5,6,7,78))

	a,b := 3,4
	swap(a,b)
	fmt.Println(a,b)

}
