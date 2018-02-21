package main

import (
	"demo/crawler/engine"
	"demo/crawler/scheduler"
	"demo/crawler/zhenai/parser"
)

const (
	SeedUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	e := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 10}

	//e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 10}

	e.Run(engine.Request{Url: SeedUrl, ParserFunc: parser.ParseCityList})
}
