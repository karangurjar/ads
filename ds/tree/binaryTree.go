package tree

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
	if node.LeftNode == nil {
		node.LeftNode = NewNode(ele)
		return
	}

	if node.RightNode == nil {
		node.RightNode = NewNode(ele)
	}

	InsertRecursive(node.LeftNode, ele)
	InsertRecursive(node.RightNode, ele)
}
