package main

import (
	"fmt"
	"strings"
)

func findCar(contents [][][]byte) string {
	for _, m := range contents {
		for _, n := range m {
			str := string(n)
			if strings.Contains(str, "车") {
				return str
			}
		}
	}
	return "not found "
}

func findHouse(contents [][][]byte) string {
	for _, m := range contents {
		for _, n := range m {
			str := string(n)
			if strings.ContainsAny(str, "住|房") {
				return str
			}
		}
	}
	return "not found "
}

func main() {
	str := "已购"


	bytes := [][][]byte{
		{[]byte(str), []byte("已购车"), []byte("未购车")},
		{[]byte(""), []byte(""), []byte("已购房")},
		{[]byte("大专"), []byte("本科"), []byte("高中"), []byte("初中")},
	}

	cars := [][][]byte{
		{[]byte("车"),[]byte("车")},
		{},
		{},
	}

	car := findCar(bytes)
	car1 := findCar(cars)
	fmt.Println("condition about car : ", car)
	fmt.Println("condition about car : ", car1)

	fmt.Println("condition about house : ",findHouse(bytes))
}
