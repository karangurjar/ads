package tree

import "fmt"

type BinaryTree struct {
	Root *Node
}

//Insert will insert the new node in level order
func (bt *BinaryTree) Insert(ele int) {
	if bt.Root == nil {
		bt.Root = NewNode(ele)
		return
	}

	queue := make([]*Node, 0)
	currentNode := bt.Root
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		deletedNode := queue[0]
		if deletedNode.LeftNode == nil {
			deletedNode.LeftNode = NewNode(ele)
			return
		}
		if deletedNode.RightNode == nil {
			deletedNode.RightNode = NewNode(ele)
			return
		}

		//insert left and right node to queue if node is not empty
		queue = queue[1:]
		queue = append(queue, deletedNode.LeftNode, deletedNode.RightNode)
	}
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
