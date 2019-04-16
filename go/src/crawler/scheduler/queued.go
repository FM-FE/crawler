package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	// we can use this struct to collect workers and requests
	// after that make a queue to store and configure them
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request  // also requestQ := []engine.Request{}
		var workerQ []chan engine.Request // also workerQ := []chan engine.Request{}

		// why those two var cannot be here
		// because every time when we start a new loop
		// activeWorker has to be nil(also means it has to be redefine since a new loop begin)
		// or the activeWorker case of select will be start
		// but actually there is no new activeWorker or activeRequest

		//var activeRequest engine.Request
		//var activeWorker chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				// got a new request, send r (request) to w, which w? we don't know, so need worker queue
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// got a new worker, send next_request to w (worker), which r? we don't know, so need request queue
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}