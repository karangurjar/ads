package tree

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

func PreOrderTraversal(root *Node) []int {
	nodesData := make([]int, 0)
	if root == nil {
		return nodesData
	}

	preOrder(root, &nodesData)

	return nodesData
}

func preOrder(root *Node, nodesData *[]int) {
	if root == nil {
		return
	}

	*nodesData = append(*nodesData, root.Data)
	preOrder(root.LeftNode, nodesData)
	preOrder(root.RightNode, nodesData)
}
