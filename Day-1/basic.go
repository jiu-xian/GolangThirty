//本节练习变量、常量、数组、切片、映射、结构体的定义和使用，以及if/for/switch/等控制语句的使用

package main

import (
	"fmt"
	"math"
)

//定义和使用变量和常量

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	//定义和使用变量和常量
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

	testArray()
	fmt.Println(arr1)
	fmt.Println("-----------------------------------")

}
