package main

import (
	"fmt"
)

func main ()  {

	var i rune

	for x := 0; x < 64; x++ {
		for i = 'a'; i <= 'z'; i++ {
			fmt.Printf("%v\n", string(i))
		}
	}

}





