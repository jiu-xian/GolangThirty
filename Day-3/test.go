package main

import (
	"fmt"
	"learngo/GolangThirty/Day-3/infra"
)

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return infra.Retriever{}

}

func main() {
	fmt.Println("第三天的练习：")
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
