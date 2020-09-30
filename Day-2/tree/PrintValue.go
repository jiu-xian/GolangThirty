package tree

import "fmt"

func (node Node) Print() {
	fmt.Println(node.Value)
	fmt.Println()
}
