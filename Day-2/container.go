//本节针对golang的内建容器进行练习，包括数组、切片、map、结构体

package main

import (
	"fmt"
	"learngo/GolangThirty/Day-2/arrays"
	"learngo/GolangThirty/Day-2/maps"
	"learngo/GolangThirty/Day-2/slices"
	"learngo/GolangThirty/Day-2/stringop"
	"learngo/GolangThirty/Day-2/structs"
)

func main() {
	fmt.Println("Golang数据结构练习：")

	arrays.Test()

	slices.Test()

	maps.Test()

	structs.Test()

	stringop.Test()

}
