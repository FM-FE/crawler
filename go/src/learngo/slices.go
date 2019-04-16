package main

import "fmt"

/*
slice的定义： s []int
不指定长度，只确定类型
 */


func updateSlice(s []int)  {
	s[0] = 100
}

func main() {
	arr := [...] int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]",arr[2:6])//arr[2:6]是一个arr的视图
	fmt.Println("arr[:6]",arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =",arr[2:])
	s2 := arr[:]
	fmt.Println("s2 =",s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]//reslice：对切片再次切片，在原先切片的基础上继续切
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n",
		s1, len(s1), cap(s1) )
	s2 = s1[3:5]//本来s1的下标到3就结束，
				// 这里s2要求从s1下标的第三位开始取取到第四位，
				// s2作为arr的切片依然保留后面的元素只是不显示
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n",
		s2, len(s2), cap(s2) )

	//slice不可以向前扩展，只能向后扩展

	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)//s4 and s5 no longer view arr;
							  // they create a new array

	fmt.Println("s3, s4, s5 =",s3,s4,s5)
	fmt.Println(arr)
}
