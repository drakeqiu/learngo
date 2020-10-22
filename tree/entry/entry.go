package main

import (
	"fmt"
	"github.com/drakeqiu/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := &myTreeNode{myNode.node.Left}
	right := &myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	/*	nodes := []tree.TreeNode{
			{Value: 3},
			{},
			{6, nil, &root},
		}
		fmt.Println(nodes)*/

	/*root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()*/

	/*root.Print()
	root.SetValue(100)

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var pRoot1 *tree.TreeNode
	pRoot1.SetValue(200)
	pRoot1 = &root
	pRoot1.SetValue(300)
	pRoot1.Print()*/

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
