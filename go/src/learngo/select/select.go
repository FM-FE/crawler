package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int)  {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id , n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	i := 0
	out := make(chan  int)
	go func() {
		for  {
			out <- i
			i++
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
		}
	}()
	return out
}

func main() {
	var c1, c2 chan int
	c1, c2 = generator(), generator()
	worker := createWorker(0)
	n := 0

	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	var values [] int
	for  {

		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <- time.After(800*time.Millisecond):
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("queue len is :",len(values))
		case <- tm:
			fmt.Println("end")
			return
		}

	}
}
