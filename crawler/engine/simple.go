package engine

import (
	"learngo/GolangThirty/crawler/fetcher"
	"log"
)

//定义一个简单的引擎，它是一个结构体
type SimpleEngine struct{}

//定义该引擎的执行函数
func (e SimpleEngine) Run(seeds ...Request) { //向run函数传入种子参数，该参数是一个不定数量的request结构体
	var requests []Request //声明一个以request结构体为元素的任务列表requests
	for _, r := range seeds {
		requests = append(requests, r) //将种子参数中的任务添加到任务列表requests中
	}
	for len(requests) > 0 { //循环执行任务列表中的任务
		r := requests[0] //提取任务列表中的第一个待执行任务
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...) //将解析结果中所包含的请求信息添加到任务列表中作为新的任务

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item) //将解析结果中的items打印出来作为日志
		}
	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url) //打印该任务的url链接作为日志

	body, err := fetcher.Fetch(r.Url) //使用抓取器中的抓取函数获取该任务中的链接所对应的body
	if err != nil {
		log.Printf("Fetcher : Error "+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err //如果获取body出错，返回出错信息
	}
	return r.ParseFunc(body), nil //对该任务链接所返回的body执行解析函数并将结果返回给一个名为parserresult的变量
}
