package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 5
var ss = "kkk" //package var not

var (
	bb = 6
	cc = true
)


func variablezero() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableinit(){
	var a int = 3
	var s string = "abc"
	fmt.Println(a,s)
}

func variabletypedecution() {
	var a,s,b,c = 1,"abs",4,true
	fmt.Println(a,s,c,b)
}

func variableshorter() {
	a,b,s,c := 1,3,"abc",true
	b = 5
	fmt.Println(a,b,s,c)
}

func euler(){
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

func triangle() {
	a, b := 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a+b*b)))
	return c
}

const filename = "abc.txt"
const (
	sss = "abc"
	bbb = 98
)

func consts()  {
	const a,b = 3,4
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c,sss,bbb)
}

func enums()  {
	const(
		cpp = iota
		java
		python
		golang
		)

	const(
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
		)

	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("hello world")
	variablezero()
	variableinit()
	variabletypedecution()
	variableshorter()
	fmt.Println(aa,ss,bb,cc)
	euler()
	triangle()
	consts()
	enums()
}

