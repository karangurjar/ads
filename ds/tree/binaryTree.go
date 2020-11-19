package tree

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(ele int) {
	if bt.Root == nil {
		bt.Root = NewNode(ele)
		return
	}

	//continue to iterate and insert
}
