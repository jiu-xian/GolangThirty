//本节练习变量、常量、数组、切片、映射、结构体的定义和使用，以及if/for/switch/等控制语句的使用

package main

import "fmt"

//定义和使用变量和常量
var a string
var b int = 2
var c, d = "hello", "world"

const g = 10

func testVariable() {
	//定义和使用变量和常量
	e, f := 3, 5 //此类语句只能在函数内使用，仅作为局部变量
	fmt.Println("打印初始变量:", a, b, c, d, e, f, g)

	a, c, d = "hello", "golang", "world"
	b, e, f = 5, 3, 9
	//g = 15	errors：此处不能再给g赋值，因为g是常量
	fmt.Println("打印赋值后的变量:", a, b, c, d, e, f, g)

	l := &a //将变量a的地址传递给指针l
	fmt.Println("打印变量的指针：")
	fmt.Println(l, *l, &b) //指针l的值（*l）

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
	testVariable()
	fmt.Println("打印赋值后的变量:", a, b, c, d, g)
	fmt.Println("------------------------------------")

	testArray()
	fmt.Println(arr1)
	fmt.Println("-----------------------------------")

}
