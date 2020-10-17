package main

import (
	"learngo/GolangThirty/crawler/engine"

	"learngo/GolangThirty/crawler/zhenai/parser"
)

func main() {
	//执行engine包的Run程序，向其输入一个request结构体作为参数，并指定该结构体的url链接和解析函数
	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList, //该解析函数为城市列表解析器中的解析函数
	})
}
