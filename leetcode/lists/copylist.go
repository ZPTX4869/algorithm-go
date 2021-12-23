package lists

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	d1 := &Node{Next: head}
	d2 := &Node{}
	nodeMap := make(map[*Node]*Node, 0)

	d1.Next = head
	p1, p2 := d1, d2
	for p1.Next != nil {
		p2.Next = &Node{
			Val:  p1.Next.Val,
			Next: nil,
		}
		nodeMap[p1.Next] = p2.Next

		p1, p2 = p1.Next, p2.Next
	}

	p1, p2 = d1, d2
	for p1.Next != nil {
		p2.Next.Random = nodeMap[p1.Next.Random]

		p1, p2 = p1.Next, p2.Next
	}

	return d2.Next
}
