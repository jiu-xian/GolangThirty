package structs

import (
	"fmt"
	"learngo/GolangThirty/Day-2/tree"
)

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

func Test() {
	fmt.Println("StructTestResult:")
	structTest()
	fmt.Println("--------------------------------")
}
