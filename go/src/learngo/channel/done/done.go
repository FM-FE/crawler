package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	done func()
	//wg *sync.WaitGroup
	//done chan bool
}

func doWork(id int, c chan int, done func()) {
	for  {
		for n := range c { // range is for traversing channel, to detect the input of channel continuously
			fmt.Printf("id is %d, recived %c\n", id, n)
			done()
			//go func() { done <- true }()
		}
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	//c := make(chan int)
	//go worker(id, c)
	//return c
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {
	wg := sync.WaitGroup{}
	var channels [10] worker
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i, &wg)
	}


	wg.Add(20)

	for i, channel := range channels {
		channel.in <- 'a' + i
	}

	for i, channel := range channels{
		channel.in <- 'A' + i
	}

	wg.Wait()

	//time.Sleep(time.Millisecond)
}

//func bufferedChannel()  {
//	c := make(chan int,3)
//	go worker(10, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	time.Sleep(time.Millisecond)
//}
//
//func channelClose()  {
//	c := make(chan int,3)
//	go worker(10, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	close(c)
//	time.Sleep(time.Millisecond)
//}

func main() {
	chanDemo()
}
