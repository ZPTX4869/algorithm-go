package list
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodes := make(map[*Node]*Node)

	var deepCopy func(curr *Node) *Node
	deepCopy = func(curr *Node) *Node {
		if curr == nil {
			return nil
		}

		if nodes[curr] != nil {
			return nodes[curr]
		}

		newNode := &Node{Val: curr.Val}
		nodes[curr] = newNode

		newNode.Next = deepCopy(curr.Next)
		newNode.Random = deepCopy(curr.Random)

		return newNode
	}

	return deepCopy(head)
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodes := make(map[*Node]*Node)
	newHead := &Node{Val: head.Val}
	nodes[head] = newHead

	p1, p2 := head, newHead
	for p1.Next != nil {
		p2.Next = &Node{Val: p1.Next.Val}
		nodes[p1.Next] = p2.Next
		p1, p2 = p1.Next, p2.Next
	}

	p1, p2 = head, newHead
	for p1 != nil {
		p2.Random = nodes[p1.Random]
		p1 = p1.Next
		p2 = p2.Next
	}

	return newHead
}
