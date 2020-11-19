package tree

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(ele int) {
	if bt.Root == nil {
		bt.Root = &Node{
			Data:      ele,
			LeftNode:  nil,
			RightNode: nil,
		}

		return
	}

	//continue to iterate and insert
}
