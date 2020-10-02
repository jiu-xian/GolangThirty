package forloop

import (
	"fmt"
	"strconv"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//循环
func adds(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
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

func forever(s string) {
	for {
		fmt.Println(s)
	}
}

func Test() {
	fmt.Println("ForloopTestResult:")
	fmt.Println(adds(100))
	//forever("abc")		这是一个无限循环，暂不执行
	fmt.Println(
		convertToBin(13),
		convertToBin(21),
		convertToBin(13729),
	)
	fmt.Println("-----------------------------------")
}
