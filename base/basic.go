package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kuuyee"

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableTypeDeduction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func euler()  {
	c:=3+4i
	fmt.Println(cmplx.Abs(c));
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi))
}

func triangle()  {
	a,b := 3,4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

func enums()  {
	const(
		cpp=iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp,javascript,python,golang)
}

func main() {
	fmt.Println("Hello kuuyee!")
	variableZeroValue()
	variableTypeDeduction()
	euler()
	triangle()
	enums()
}
