package test

import (
	"fmt"
	"learngo/GolangThirty/Day-3/fib"
	"testing"
)

//以下是测试的代码练习
func TestFib(t *testing.T) {
	tests := []struct{ a, b int }{
		{10, 10},
		{20, 20},
		{100, 100},
		{1000, 1000},
	}
	for _, tt := range tests {
		if actual := len(fib.CreateFib(tt.a)); actual != tt.b {
			t.Errorf("CreateFib(%d )"+"Got %d ; expected %d", tt.a, actual, tt.b)
		}
	}
}

//以下是代码示例的代码练习
func ExampleCreateFib() {
	l1 := len(fib.CreateFib(10))
	l2 := len(fib.CreateFib(20))
	l3 := len(fib.CreateFib(30))
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)

	//Output:
	//10
	//20
	//30
}
