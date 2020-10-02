package consts

import (
	"fmt"
	"math"
)

//3:常量和枚举的方法：
func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4 //此处未定义a，b的类型
	var c int
	c = int(math.Sqrt(a*a + b*b)) //此处根据需要将a，b类型以float来处理
	fmt.Println(filename, a, b, c)
}

//使用iota进行递增枚举的示例：
func enums() {
	const (
		golang = 1
		python = 2
		cpp    = 3
		java   = 4
	)
	const (
		a = iota
		b
		c
		d
	)
	const (
		e = 1 << (10 * iota)
		f
		_
		h
		i
	)
	fmt.Println(golang, python, cpp, java)
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, h, i)
}

func Test() {
	fmt.Println("ConstsTestResult:")
	consts()
	enums()
	fmt.Println("-----------------------------------")
}
