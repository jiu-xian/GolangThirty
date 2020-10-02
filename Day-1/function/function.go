package function

import (
	"fmt"
	"math"
)

//函数的定义和使用
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("不能除以0")
	}
	return a / b, nil
}

//此处需体会函数作为参数进行传递的用法
func apply(op func(a, b int) int, a, b int) int {
	return op(a, b)

}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Test() {
	fmt.Println("FunctionTestResult:")
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 4))

	fmt.Println(apply(pow, 3, 4))
	fmt.Println("-----------------------------------")
}
