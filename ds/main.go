package main

import (
	"fmt"

	"github.com/karsingh991/ads/ds/tree"
)

func main() {
	var bt tree.BinaryTree
	bt.Insert(10)
	fmt.Printf("%+v", bt)
}
