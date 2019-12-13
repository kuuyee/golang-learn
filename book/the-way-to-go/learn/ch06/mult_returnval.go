package main

import "fmt"

func main() {
	fmt.Println(noNameFn(9, 5))
	fmt.Println(nameFn(7, 2))

}

func noNameFn(a, b int) (int, int, int) {
	sum := a + b
	pro := a * b
	dif := a - b

	return sum, pro, dif
}

func nameFn(a, b int) (sum int, pro int, dif int) {
	sum = a + b
	pro = a * b
	dif = a - b

	return
}
