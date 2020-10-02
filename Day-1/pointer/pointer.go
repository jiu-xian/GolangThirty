package pointer

import "fmt"

//指针的概念和用法
func pointerTest() {
	var a int = 2
	var pa *int = &a //pa是一个整数的指针，指针的值是变量a的地址（&a）
	*pa = 3          //指针pa所指向的值（*pa）重新赋值为3
	fmt.Println(a)
	fmt.Println(pa)
}

func swap1(a, b int) {
	a, b = b, a //此处由于是值传递，函数再传入a，b时拷贝了一份a，b的值，所以原始的a，b的值并未改变
}

func swap2(a, b *int) {
	*a, *b = *b, *a //此处a，b是指针，将指针所指向的值进行互换赋值，可以改变原始a，b的值
}

func swap3(a, b int) (int, int) {
	return b, a //此处通过将函数运算的结果返回出去，可以改变原始a，b的值
}

func swapTest() {
	a, b := 3, 4
	swap1(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)

	a, b = swap3(a, b)
	fmt.Println(a, b)

}

func Test() {
	fmt.Println("PointerTestResult:")
	pointerTest()
	swapTest()
	fmt.Println("--------------------------")
}
