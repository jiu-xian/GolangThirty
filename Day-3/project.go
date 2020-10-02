package main

import (
	"fmt"
	"learngo/GolangThirty/Day-3/fib"
	"learngo/GolangThirty/Day-3/testing"
	"time"
)

//声明一个接口类型
type retriever interface {
	Get(string) string
}

//通过这种方式轻松切换接口指向
func getRetriever() retriever {
	return testing.Retriever{}

}

//defer的用法：
func tryDefer() {
	sum := 0
	for i := 0; i < 30; i++ {
		sum += i
		defer fmt.Println(sum)
	}

}

//recover的用法：
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

//goroutine用法：
func goPrint() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("hello", i)
			}
		}(i) //注意这里的匿名函数用法
	}
	time.Sleep(time.Millisecond)
}

//channel用法
func chanDemo() {
	c := make(chan int) //创建一个channel
	go func() {
		for {
			n := <-c //从channel接收值
			fmt.Println(n)
		}
	}()
	c <- 1 //向channel传递值
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

//buffered channel的用法：
func bufferedChannel() {
	c := make(chan int, 3) //channel的容量设置为3
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

//遍历channel
func worker2(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

//close channel
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
