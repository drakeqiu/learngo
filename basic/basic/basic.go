package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	//fmt.Println(
	//	cmplx.Pow(math.E, 1i * math.Pi) + 1)

	//fmt.Println(
	//	cmplx.Exp(1i * math.Pi) + 1)

	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//const filename = "abc.txt"
func consts() {
	//const filename = "abc.txt"
	//const a, b = 3, 4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	/*const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)*/

	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()

	consts()
	enums()
}
