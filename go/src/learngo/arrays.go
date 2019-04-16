package main

import "fmt"

func printArray(arr []int)  {//只有当参数为指针类型的时候，才会改变数组的值
	arr[0] = 100
	for i,v := range arr{//new;下划线省略变量
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5] int
	arr2 := [3] int{1,3,5}
	arr3 := [...] int{2,4,6,8,10}

	var grid [4][5] int

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	fmt.Println("arr1:")
	printArray(arr1[:])
	fmt.Println("arr3:")
	printArray(arr3[:])

	fmt.Println(arr1,arr3)//数组是值类型，所以在调用的时候都需要做一份拷贝

	//printArray(arr2): arr2是 [3]int ；arr3是 [5]int；go语言认为是不一样的类型



}
