package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	/*root.print()
	root.setValue(100)

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var pRoot1 *treeNode
	pRoot1.setValue(200)
	pRoot1 = &root
	pRoot1.setValue(300)
	pRoot1.print()*/

	root.traverse()
}
