package stringop

import (
	"fmt"
	"unicode/utf8"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//字符串处理

func stringOperate() {
	s := "yes我爱饭米粒！"
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { //ch 就是一个rune类型
		fmt.Printf("(%d %X  ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)  ", i, ch)
	}
}

func Test() {
	fmt.Println("StringopTestResult:")
	stringOperate()
	fmt.Println("----------------------------------")
}
