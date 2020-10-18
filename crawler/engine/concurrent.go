package engine

import (
	"learngo/GolangThirty/crawler/model"
	"log"
)

type ConcurrentEngine struct { //定义一个并发版的引擎
	Scheduler Scheduler //引擎中的队列是一个名为队列的接口
	WorkCount int       //引擎中的并发数是一个整数
}

type Scheduler interface { //定义个名为队列的接口，该接口包含三个函数和一个子接口
	Submit(Request)           //此函数用于将一个请求提交到channel中
	WorkerChan() chan Request //此函数返回一个类型为request的channel
	Run()                     //此函数为队列调度器的执行程序
	ReadyNotifier             //这是一个子接口，用于
}

type ReadyNotifier interface { //定义一个名为readynotifier的子接口
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) { //并发引擎的执行程序

	out := make(chan ParseResult)
	e.Scheduler.Run() //启动调度器

	for i := 0; i < e.WorkCount; i++ { //创建一些列并发的工作任务
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r) //将任务种子中的请求任务提交给请求通道
	}

	profileCount := 0
	for {
		result := <-out //将工作结果打印计数
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok { //对item进行强制类型转换看看是否能成功来判断item是不是一个profile的类型
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}
		}
		//Url dedup 去重

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request) //将新请求提交到请求通道中
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) { //创建一个创建工作的函数
	go func() {
		for {
			ready.WorkerReady(in) //开始准备工作，接收到一个请求通道传入，并将其加入工作通道
			//告诉scheduler 准备好了
			request := <-in                //提取请求通道中的请求任务
			result, err := worker(request) //处理请求任务
			if err != nil {
				continue
			}
			out <- result //将处理结果传递给out通道
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
