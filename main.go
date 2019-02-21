package main

import (
	"learnGo/crawier/engine"
	"learnGo/crawier/persist"
	"learnGo/crawier/scheduler"
	"learnGo/crawier/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSave(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
