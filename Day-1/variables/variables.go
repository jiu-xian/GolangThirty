package variables

import "fmt"

//1：定义和使用变量
//1.1 变量的声明
//1.2 变量的使用范围

//声明包内变量的示例：
var (
	aa = 3
	ss = "kkk"
	bb = true
)

//声明一组默认值变量的示例：
func variableZeroValue() {
	var a int
	var b string
	var c bool
	fmt.Printf("%d %q %t\n", a, b, c)
}

//声明一组变量并进行初始赋值
func variableInitialValue() {
	var a, d int
	var b string
	var c bool
	a, d = 3, 5
	b = "hello"
	c = false
	fmt.Printf("%d %q %t %d\n", a, b, c, d)
}

//可以让编译器推测变量的类型
func variableTypeDeduction() {
	var a, b, c, d = 3, 5, "hello", true
	fmt.Printf("%d %d %s %t\n", a, b, c, d)
}

//简短变量声明，仅限于在函数体内使用
func variableShorter() {
	a, b, c, d := 3, 5, "hello", true
	b = 5
	fmt.Printf("%d %d %s %t\n", a, b, c, d)
}
func Test() {
	fmt.Println("VariableTestResult:")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
}
