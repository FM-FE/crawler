package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserRequests, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserRequests.Requests...)

		for _, item := range parserRequests.Items {

			log.Printf("Got item : %+v \n", item)
		}
	}

}

func worker(r Request) (ParserRequests, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching Url %s: %v", r.Url, err)
		return ParserRequests{}, err
	}

	return r.ParserFunc(body), nil
}
