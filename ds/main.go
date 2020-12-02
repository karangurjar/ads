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
	fmt.Printf("%+v\n", bt.Root)
	fmt.Printf("%+v\n", bt.Root.LeftNode)
	fmt.Printf("%+v\n", bt.Root.RightNode)
}
