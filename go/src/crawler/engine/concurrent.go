package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{} // ItemSaver
}

//type Scheduler interface { // simple scheduler
//	Submit(Request)
//	ConfigureMasterWorkerChan(chan Request)
//	WorkerReady(chan Request)
//	Run()
//}

//type Scheduler interface { // queue scheduler
//	Submit(Request)
//	WorkerReady(chan Request)
//	Run()
//}

type Scheduler interface {
	// code refactoring
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

//func createWorker(in chan Request, out chan ParserRequests) { // simple scheduler
//	go func() {
//		for {
//			request := <- in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}

//func createWorker(out chan ParserRequests, scheduler Scheduler) { // queue scheduler
//	go func() {
//      in := make(chan Request)
//		for {
//			// tell scheduler, i am ready
//			scheduler.WorkerReady(in)
//			request := <- in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}

func createWorker(in chan Request, out chan ParserRequests, scheduler Scheduler) { // code refactoring
	go func() {
		for {
			// tell scheduler, i am ready
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// simple scheduler
	//in := make(chan Request)
	//out := make(chan ParserRequests)
	//e.Scheduler.ConfigureMasterWorkerChan(in)

	// queue scheduler
	out := make(chan ParserRequests)
	e.Scheduler.Run()

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i := 1; i <= e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}
