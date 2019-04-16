package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// <a href="http://www.zhenai.com/zhenghun/aba" data-v-0c63b635>阿坝</a>

//func getCityList(contents []byte) {
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)
//	// care about the [^>] and [^<], both of those define the limit of regexp
//	matches := re.FindAllSubmatch(contents, -1)
//
//	for _, m := range matches {
//		fmt.Printf("city is %s, URL is %s \n", m[2], m[1])
//	}
//	fmt.Println("length of matches is : ", len(matches))
//}

func main() {
	// &scheduler.SimpleScheduler{} is a pointer receiver
	// so we have to define it first
	// or compile cannot find the address
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})

}
