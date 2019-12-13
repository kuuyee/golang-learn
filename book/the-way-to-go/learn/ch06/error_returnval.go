package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySqrt(-4.0))
	fmt.Println(MySqrtNoName(9))
}

func MySqrt(fs float64) (float64, error) {
	if fs < 0 {
		return float64(math.NaN()), errors.New("小于0不能计算平方根！")
	}
	fsSqrt := math.Sqrt(fs)
	return fsSqrt, nil
}

func MySqrtNoName(fs float64) (fsSqrt float64) {
	fsSqrt = math.Sqrt(fs)
	return
}
