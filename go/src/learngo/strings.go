package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我爱六四七!" // chinese use three byte to store, so called UTF-8
	fmt.Printf("%X\n",s)

	for _,v := range []byte(s) {
		fmt.Printf("%X  ",v)
	}
	fmt.Println()

	for i, v := range s { // ch is a rune
		fmt.Printf("(%d, %X)",i, v)
	}
	fmt.Println()

	fmt.Println("Rune count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c  ",ch)
	}
	fmt.Println()

	for i, v := range []rune(s) { // change s to a rune slice; that does not mean the old s, rune get a new  to store
		fmt.Printf("(%d, %c)", i, v)
	}
	fmt.Println()

}
