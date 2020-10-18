package scheduler

import "learngo/GolangThirty/crawler/engine"

type QueuedScheduler struct { //定义一个队列调度器
	requestChan chan engine.Request      //请求channel，类型为request
	workerChan  chan chan engine.Request //任务channel，类型为类型为request的channel
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request { //实现接口的WorkerChan功能，创建一个新的类型为request的channel
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) { //实现接口的Submit功能，将接收到的r存入到请求channel中去
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) { //实现接口的WorkerReady功能，将接收到的request channel 传递给 workerchannel
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() { //队列调度器的执行程序
	s.workerChan = make(chan chan engine.Request) //创建队列调度器中的两个通道，工作通道和请求通道
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request     //定义请求队列，队列中是request
		var workerQ []chan engine.Request //定义工作通道队列，队列中是存有可接收request的channel
		for {
			var activeRequest engine.Request //定义活跃状态的请求任务和工作通道
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 { //工作队列中有闲置工作通道，且请求队列中有未处理请求时，
				activeWorker = workerQ[0]   //提取工作通道
				activeRequest = requestQ[0] //提取任务
			}

			select {
			case r := <-s.requestChan: //当接收到新的请求是，将其添加到队列中
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: //当有新的空闲工作通道被释放时，将其添加到工作通道队列
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //当提取出的请求任务成功传入工作通道后，更新请求队列和工作通道队列
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
