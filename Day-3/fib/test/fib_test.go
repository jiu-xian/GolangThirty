package test

import (
	"learngo/GolangThirty/Day-3/fib"
	"testing"
)

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
