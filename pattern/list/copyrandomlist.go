package list

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

//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return nil
//	}
//
//	for curr := head; curr != nil; {
//		dupNode := &Node{curr.Val, curr.Next, nil}
//		dupNode.Next = curr.Next
//		curr.Next = dupNode
//		curr = dupNode.Next
//	}
//
//	for curr := head; curr != nil; {
//		if curr.Random != nil {
//			curr.Next.Random = curr.Random.Next
//		}
//		curr = curr.Next.Next
//	}
//
//	newHead := head.Next
//	for curr := head; curr != nil; curr = curr.Next {
//		newNode := curr.Next
//		curr.Next = newNode.Next
//
//		if curr.Next == nil {
//			break
//		}
//
//		newNode.Next = curr.Next.Next
//	}
//
//	return newHead
//}

//func copyRandomList(head *Node) *Node {
//	nodeMap := make(map[*Node]*Node)
//
//	var deepCopy func(head *Node) *Node
//	deepCopy = func(head *Node) *Node {
//		if head == nil {
//			return nil
//		}
//
//		if createdNode, ok := nodeMap[head]; ok {
//			return createdNode
//		}
//
//		newNode := &Node{Val: head.Val}
//		nodeMap[head] = newNode
//		newNode.Next = deepCopy(head.Next)
//		newNode.Random = deepCopy(head.Random)
//
//		return newNode
//	}
//
//	return deepCopy(head)
//}
