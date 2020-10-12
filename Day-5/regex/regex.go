package main

import (
	"fmt"
	"regexp"
)

const text = "My Email is alex@gmail.com"
const text2 = "My Email is alex@gmail.com.cn"
const text3 = `
My Email is alex@gmai.com
My Email is Bob@gmai.com.cn
My Email is Char@def.org
`

func regexTest1() {
	re, err := regexp.Compile("alex@gmail.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(re)
}

func regexTest2() {
	re := regexp.MustCompile("alex@gmail.com")
	match := re.FindString(text)
	fmt.Println(match)
}

func regexTest3() {
	re := regexp.MustCompile(`.+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)
}

func regexTest4() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	match := re.FindString(text2)
	fmt.Println(match)
}

func regexTest5() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	match := re.FindAllString(text3, -1)
	fmt.Println(match)
}

func regexTest6() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text3, -1)
	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}

func main() {
	regexTest1()
	regexTest2()
	regexTest3()
	regexTest4()
	regexTest5()
	regexTest6()
}
