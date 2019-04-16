package main

import (
	"fmt"
	"reflect"
)

type t struct {
	a string
	b string
	c string
}

func newtest(a, b, c string) ( *t) {
	return &t{
		a: a,
		b: b,
		c: c,
	}
}

func (t t)test()  {
	fmt.Printf("a is string: %v", reflect.TypeOf(t.a))
	fmt.Printf("b is string: %v", reflect.TypeOf(t.b))
	fmt.Printf("c is string: %v", reflect.TypeOf(t.c))
}

func main() {
	t := newtest("1", "2", "3")
	t.test()
}
