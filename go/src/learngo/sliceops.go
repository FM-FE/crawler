package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v, len = %d, cap = %d\n", s, len(s), cap(s))
}

func main() {
	var s []int//zero value for slice is nil

	fmt.Println("Creating slice")
	for i := 0; i < 100; i++ {
		s = append(s, 2 * i +1)
	}
	fmt.Println(s)

	s1 := []int {2,4,6,8} // the way to create a slice, use []int to do
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10,32)//the first element is len, the second is cap
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)//s1(source) copy to s2(destination)
	printSlice(s2)

	fmt.Println("deleting slice")
	s2 = append(s2[:3], s2[4:] ...)
	printSlice(s2)

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Printf("front = %d, slice = %v, tail = %d",
		front, s2, tail)

}
