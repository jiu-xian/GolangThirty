//本节练习变量、常量、数组、切片、映射、结构体的定义和使用，以及if/for/switch/等控制语句的使用

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//1：定义和使用变量

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var b string
	var c bool
	fmt.Printf("%d %q %t\n", a, b, c)
	fmt.Println("-----------------------------------")

}

func variableInitialValue() {
	var a, d int
	var b string
	var c bool
	a, d = 3, 5
	b = "hello"
	c = false
	fmt.Printf("%d %q %t %d\n", a, b, c, d)
	fmt.Println("-----------------------------------")

}

func variableTypeDeduction() {
	var a, b, c, d = 3, 5, "hello", true
	fmt.Printf("%d %d %s %t\n", a, b, c, d)
	fmt.Println("-----------------------------------")
}

func variableShorter() {
	a, b, c, d := 3, 5, "hello", true
	b = 5
	fmt.Printf("%d %d %s %t\n", a, b, c, d)
	fmt.Println("-----------------------------------")
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//3:常量和枚举
func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4 //此处未定义a，b的类型
	var c int
	c = int(math.Sqrt(a*a + b*b)) //此处根据需要将a，b类型以float来处理
	fmt.Println(filename, a, b, c)
	fmt.Println("-----------------------------------")
}

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
	fmt.Println("-----------------------------------")
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//4-条件语句
func iftest1(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func iftest2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//if代码块外无法再打印contents，因为contents变量属于if区块内的变量
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsurpported operator: " + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//循环
func adds() int {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	var result string = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//函数
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("不能除以0")
	}
	return a / b, nil
}

func apply(op func(a, b int) int, a, b int) int {
	return op(a, b)

}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//指针
func pointerTest() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	fmt.Println(pa)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

	consts()
	enums()

	fmt.Println(iftest1(120))
	fmt.Println(iftest1(-30))
	fmt.Println(iftest1(97))
	iftest2()
	fmt.Println(eval(3, 5, "+"))
	fmt.Println(eval(3, 5, "-"))
	fmt.Println(eval(3, 5, "*"))
	fmt.Println(eval(3, 5, "/"))
	//fmt.Println(eval(3, 5, "a"))

	fmt.Println(
		grade(53),
		grade(76),
		grade(98),
	)

	fmt.Println(adds())
	//forever()		这是一个无限循环，暂不执行
	fmt.Println(
		convertToBin(13),
		convertToBin(21),
		convertToBin(13729),
	)
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 4))

	fmt.Println(apply(pow, 3, 4))

	pointerTest()

}
