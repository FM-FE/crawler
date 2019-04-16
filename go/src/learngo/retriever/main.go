package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const  url = "http://www.fm-fe.com"

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents":"another faked site",
	})
	return s.Get(url)
}


func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name":"fm-fe",
		"course":"golang",
	})
}


func main() { // go语言所有的变量都是值类型，没有指针类型，
			  // 也就是说所有的数据无论定义了指针还是值，里面都存有数据，指针不是只有地址
	var r Retriever

	mockretriever := mock.Retriever{"this is fake site"}
	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockretriever))
	fmt.Println()

	inspect(&mockretriever)

	//Type assertion
	//mockRetriever := mock.Retriever{"test"}
	//mockeRetriever := r.(*mock.Retriever)
	//fmt.Println(mockeRetriever)

	r = &real2.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(r)

	//Type assertion
	//realRetriever := r.(*real2.Retriever)
	//fmt.Println(realRetriever.UserAgent)

	//fmt.Println(download(r))
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting: ",r)
	fmt.Printf("> %T %v\n",r, r)
	fmt.Print("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
	fmt.Println()
}