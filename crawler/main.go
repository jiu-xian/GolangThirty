package main

import (
	"learngo/GolangThirty/crawler/engine"
	"learngo/GolangThirty/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
