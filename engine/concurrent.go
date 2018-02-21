package engine

import (
	"log"
)

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	cityTotalCount := 0
	profileTotalCount := 0
	profileDetailsTotalCount := 0

	for {
		result := <-out

		// handle data
		for _, item := range result.Items {
			if result.ItemType == CityRequest {
				log.Printf("[%04d]City: %v", cityTotalCount+1, item)
				cityTotalCount++
			} else if result.ItemType == ProfileRequest {
				log.Printf("[%05d]Profile name : %v", profileTotalCount+1, item)
				profileTotalCount++
			} else if result.ItemType == ProfileDetails {
				log.Printf("[%05d]Profile details : %v", profileDetailsTotalCount+1, item)
				profileDetailsTotalCount++
			}
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
