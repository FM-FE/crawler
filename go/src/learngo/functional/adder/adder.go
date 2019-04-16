package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		fmt.Printf("> v = %d",v)
		fmt.Println()
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		//fmt.Print(">  ")
		//fmt.Println(base + v, adder2(base + v))
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder2(0)
	var s int
	for i := 0; i < 10; i++ {
		s, a = a(i) // when we change a to _, the output also change
		fmt.Printf("0+1+ ... %d = %d   \n", i, s)
		//fmt.Println( i, s, a)
	}
}
