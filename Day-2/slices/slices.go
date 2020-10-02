package slices

import (
	"fmt"
)

//切片
func sliceTest() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:5]
	s2 := arr[:]
	fmt.Println("arr[:5] = ", s1)
	fmt.Println("arr[:] = ", s2)
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:] = ", arr[2:])

	//slice的长度和容量
	s1 = s1[2:5]
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	s1 = s1[3:5]
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
}

//slice的扩展
func sliceOperat() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[3:5]
	fmt.Println(s1)

	s2 := append(s1, 10)
	s3 := append(s2, 11)
	s4 := append(s3, 12)
	s5 := append(s4, 13)
	fmt.Println("s2 s3 s4 s5 = ", s2, s3, s4, s5)
	fmt.Println(arr)
}

func printSlice(s []int) {
	fmt.Println(s, ", len = ", len(s), ", cap = ", cap(s))
}

func creatSlice() {
	var s1 = []int{} //slice的zerovalue是nil
	s2 := []int{1, 2, 3, 4, 5, 6}
	s3 := make([]int, 10)     //确定slice的len
	s4 := make([]int, 16, 32) //确定slice的len和cap
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	printSlice(s4)

	for i := 0; i < 100; i++ {
		s1 = append(s1, 2*i+1)
	}
	printSlice(s1) //slice增加元素时，如果cap不足，会额外增加容量备用

	copy(s3, s2) //把s2的元素拷贝到s3中去
	printSlice(s3)

	s3 = append(s3[:3], s3[4:]...) //删除第3个元素，注意s3[4:]...是将s3中第4个元素开始以后的每个元素逐一添加到s3[:3]中
	printSlice(s3)

	s3 = s3[1:] //删除第一个元素
	printSlice(s3)

	tail := len(s3) - 1
	s3 = s3[:tail] //删除最后一个元素
	printSlice(s3)

}

func Test() {
	fmt.Println("SliceTestResult:")
	sliceTest()
	sliceOperat()
	creatSlice()
	fmt.Println("----------------------------------")
}
