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

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func chanDemo2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func worker2(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func closeChannel() {
	c := make(chan int, 3)
	go worker2(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
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

	//goPrint()

	chanDemo()
	chanDemo2()
	bufferedChannel()
	closeChannel()

}
