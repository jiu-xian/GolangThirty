package main

import (
	"learngo/GolangThirty/crawler/engine"
	"learngo/GolangThirty/crawler/scheduler"

	"learngo/GolangThirty/crawler/zhenai/parser"
)

func main() {
	//创建一个引擎
	e := engine.ConcurrentEngine{ //此处的concurrentengine可更换为simpleengine等其他结构体
		Scheduler: &scheduler.QueuedScheduler{}, //引擎中的调度器设置为队列调度器queued scheduler
		WorkCount: 100,                          //workcount并发数设置为100
	}

	//执行该引擎的Run程序，向其输入一个request结构体作为参数，并指定该结构体的种子url链接和解析函数
	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList, //该解析函数为城市列表解析器中的解析函数
	})

	/*
		//此处为测试单一城市页面的代码
		e.Run(engine.Request{
			Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai",
			ParseFunc: parser.ParseCity,
		})
	*/

}
