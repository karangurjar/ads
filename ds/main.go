package main

import (
	"fmt"

	"github.com/karsingh991/ads/ds/tree"
)

func main() {
	var bt tree.BinaryTree
	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(30)
	bt.Insert(40)
	bt.Insert(50)
	bt.Insert(60)
	fmt.Print(tree.PreOrderTraversal(bt.Root))
}
