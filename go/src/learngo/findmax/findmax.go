package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s)  {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start { //this line can make sure the string start at the second letter
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i // this step is give map[byte] a int, so we can get lastOccurred[ch] no matter when
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("aaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("abbbcde"))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("旋律缓慢而忧伤"))
	// this algorithm is excellent
	// rune is great, it can help you to deal with chinese or other language
}
