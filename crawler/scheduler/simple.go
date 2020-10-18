package scheduler

import "learngo/GolangThirty/crawler/engine"

type SimpleScheduler struct { //定义一个简单的调度器，它包含一个参数
	workerChan chan engine.Request //此参数是一个类型为request的channel
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request { //将参数workerchan通过接口中的WorkerChan函数传递出去
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) { //空函数，只是为了和接口对应
}

func (s *SimpleScheduler) Run() { //该队列调度器的执行函数，创建一个workerchan变量
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) { //该调度器通过workerchan不断的接收到新request
	go func() { s.workerChan <- r }() //通过goroutine避免channel拥堵卡死
}
