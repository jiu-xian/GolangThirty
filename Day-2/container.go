//本节针对golang的内建容器进行练习，包括数组、切片、map

package main

import (
	"fmt"
	"learngo/GolangThirty/Day-2/tree"
	"unicode/utf8"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//数组
var arr1 [5]int

func arrayTest() {
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	arr1 = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//切片
func sliceTest() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:5]
	s2 := arr[:]
	fmt.Println("arr[:5] = ", s1)
	fmt.Println("arr[:] = ", s2)
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:] = ", arr[2:])

	s1 = s1[2:5]
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	s1 = s1[3:5]
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
}

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

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//映射map
func mapTest() {
	m1 := map[string]string{
		"name":     "jiuxian",
		"learning": "golang",
		"feeling":  "notbad",
		"purpose":  "blockchain",
	}
	var m2 map[string]int      //此处的zerovalue是nil
	m3 := make(map[string]int) //此处的zerovalue是0

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	a := m1["name"]
	fmt.Println(a)

	b := m1["nime"]
	fmt.Println(b)

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	delete(m1, "name")

	if v, ok := m1["name"]; ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("您要找的值不存在", "m1的元素数量为： ", len(m1))
	}
}

func nonRepeating(s string) int {
	lastoccurred := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, v := range []byte(s) {
		if lasti, ok := lastoccurred[v]; ok && lasti >= start {
			start = lasti + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}

		lastoccurred[v] = i
	}
	return maxlength
}

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

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//结构体
type myTreeNode struct {
	node *tree.Node
}

func (mynode *myTreeNode) downTraverse() {
	if mynode == nil || mynode.node == nil {
		return
	}
	left := myTreeNode{mynode.node.Left}
	left.downTraverse()
	right := myTreeNode{mynode.node.Right}
	right.downTraverse()
	mynode.node.Print()
}

func structTest() {
	var root tree.Node
	root2 := tree.Node{}
	root3 := tree.Node{Value: 6, Left: nil, Right: nil}

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 6},
		{},
		{Value: 5, Left: nil, Right: &root},
	}
	fmt.Println(root, root2, root3, nodes)

	root.Print()
	root2.Print()
	root3.Print()

	root.Left.Right.SetValue(4)
	root.Left.Right.Print()
	fmt.Println("遍历此结构：")
	root.Traverse()
	myroot := myTreeNode{node: &root}
	myroot.downTraverse()
}

func main() {
	fmt.Println("Golang数据结构练习：")
	arrayTest()
	printArray(arr1)
	fmt.Println(arr1)
	printArray2(&arr1)
	fmt.Println(arr1)
	fmt.Println("----------------------------------")

	sliceTest()
	sliceOperat()
	creatSlice()
	fmt.Println("----------------------------------")

	mapTest()

	fmt.Println(nonRepeating("abcdacdebcd"))
	fmt.Println(nonRepeating("abcdefg"))
	fmt.Println(nonRepeating(""))
	fmt.Println(nonRepeating("我爱family！"))
	fmt.Println(nonRepeating("一二三三二一"))
	fmt.Println("----------------------------------")

	stringOperate()
	fmt.Println("----------------------------------")

	structTest()

}
