package arrays

import (
	"fmt"
)

//数组
var arr1 [5]int //定义一个5个元素的整数数组

func arrayTest() {
	arr2 := [3]int{1, 2, 3}          //创建一个拥有3个元素的整数数组
	arr3 := [...]int{2, 4, 6, 8, 10} //由编译器自动计算元素个数
	var grid [4][5]int               //定义一个4行5列的整数数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	arr1 = [5]int{1, 2, 3, 4, 5} //给数组arr1赋值

	for i := 0; i < len(arr1); i++ { //遍历数组
		fmt.Println(arr1[i])
	}
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) { //使用指针改变原始数组的值
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func Test() {
	fmt.Println("ArrayTestResult:")
	arrayTest()
	printArray(arr1)
	fmt.Println(arr1) //此处arr1未被改变
	printArray2(&arr1)
	fmt.Println(arr1) //此处arr1已经改变
	fmt.Println("----------------------------------")
}
