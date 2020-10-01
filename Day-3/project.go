package main

import (
	"fmt"
	"learngo/GolangThirty/Day-3/fib"
	"learngo/GolangThirty/Day-3/testing"
	"time"
)

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return testing.Retriever{}

}

func tryDefer() {
	sum := 0
	for i := 0; i < 30; i++ {
		sum += i
		defer fmt.Println(sum)
	}

}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("错误提示：", err)

		} else {
			fmt.Println("我也没辙。。。")
		}
	}()

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123)
}

func goPrint() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("hello", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("第三天的练习：")
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))

	s := fib.CreateFib(30)
	fmt.Println(s, len(s))
	tryDefer()

	tryRecover()

	goPrint()
}
