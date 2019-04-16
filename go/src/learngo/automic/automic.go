package main

import (
	"fmt"
	"sync"
	"time"
)

type automicInt struct {
	value int
	lock  sync.Mutex
}

func (a *automicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *automicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a automicInt
	done := make(chan bool) // have to use make to create a chan; if use var, it will be nil chan
	a.increment()
	go func() {
		a.increment()
		done <- true
	}()

	time.Sleep(time.Millisecond)

	fmt.Println(a.get())

	//n := <- done
	//if n{
	//	fmt.Println(a)
	//	fmt.Println(a.get())
	//}
}
