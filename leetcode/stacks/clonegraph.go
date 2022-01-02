package stacks

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	head := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}

	queue := make([]int, 0)
	for len(queue) != 0 {

	}

	return head
}
