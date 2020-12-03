package tree

import "fmt"

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(ele int) {
	if bt.Root == nil {
		bt.Root = NewNode(ele)
		return
	}
	InsertRecursive(bt.Root, ele)
}

func InsertRecursive(node *Node, ele int) {
	fmt.Println("in recursiion: ", node)
	if node.LeftNode == nil {
		node.LeftNode = NewNode(ele)
		return
	}

	if node.RightNode == nil {
		node.RightNode = NewNode(ele)
		return
	}

	InsertRecursive(node.LeftNode, ele)
	InsertRecursive(node.RightNode, ele)
}

func (bt *BinaryTree) Print() {
	fmt.Print("[")
	Traverse(bt.Root)
	fmt.Print("]")
}

func Traverse(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d,", root.Data)
	Traverse(root.LeftNode)
	Traverse(root.RightNode)
}
