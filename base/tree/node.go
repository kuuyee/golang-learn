package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}

func (node *treeNode) traverse()  {
	node.left.traverse()
	node.print()
	node.right.traverse()
}
func main() {
	var root treeNode
	fmt.Println(root)

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
}
