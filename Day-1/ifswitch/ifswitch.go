package ifswitch

import (
	"fmt"
	"io/ioutil"
)

//4-条件语句的使用
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

//switch语句的使用
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

//siwtch后可以没有表达式
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

func Test() {
	fmt.Println("IfswitchTestResult:")
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
	fmt.Println("-----------------------------------")
}
