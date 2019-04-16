package main

import (
	"fmt"
	"time"
)

func main() {
	//var a[10] int
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				// a[i]++
				// runtime.Gosched()
				fmt.Println("Hello from goroutine :",i)
			}
		}()
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}