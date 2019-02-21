package main

import (
	"learnGo/crawier/engine"
	"learnGo/crawier/scheduler"
	"learnGo/crawier/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
