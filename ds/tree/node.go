package tree

type Node struct {
	Data      int
	LeftNode  *Node
	RightNode *Node
}

func NewNode(ele int) *Node {
	return &Node{
		Data:      ele,
		LeftNode:  nil,
		RightNode: nil,
	}

}
