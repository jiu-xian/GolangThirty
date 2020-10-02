//本节练习变量、常量、以及if/for/switch/等控制语句的使用

package main //包的名称，一个目录下只能有一个包，但可以有多个同一个报名的.go文件

import (
	"learngo/GolangThirty/Day-1/consts"
	"learngo/GolangThirty/Day-1/forloop"
	"learngo/GolangThirty/Day-1/function"
	"learngo/GolangThirty/Day-1/ifswitch"
	"learngo/GolangThirty/Day-1/pointer"
	"learngo/GolangThirty/Day-1/variables"
)

func main() {
	variables.Test()

	consts.Test()

	ifswitch.Test()

	forloop.Test()

	function.Test()

	pointer.Test()

}
