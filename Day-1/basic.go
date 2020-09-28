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
func convertToBin(n int) string {
	var result string = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
var arr1 = [3]int{21, 22, 23}

func testArray() {
	//定义和使用数组
	var arr2 [3]int
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [3]string{"hello", "golang", "world"}
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	var grid [4][5]int
	fmt.Println("打印原始数组：")
	fmt.Println(arr1, arr2, arr3, arr4, arr5, grid)

	arr2 = [3]int{11, 12, 13}
	//arr3 = [4]int{1, 2, 3, 4}  errors：[5]int和[4]int是两种不同的类型，不能传递
	//arr5 = [3]int{1, 2, 3}    errors:不同类型数组不能传递值
	fmt.Println("打印传递值之后的数组：")
	fmt.Println(arr2)

	arr2[0] = 111
	arr1[0] = 222
	fmt.Println(arr1, arr2)

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

	fmt.Println(
		convertToBin(13),
		convertToBin(21),
		convertToBin(13729),
	)

	testArray()
	fmt.Println(arr1)
	fmt.Println("-----------------------------------")

}
