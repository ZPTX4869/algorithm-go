package lists

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	prev.Next = nil

	tail := reverse(slow)

	p1, p2 := head, tail
	for p1 != nil {
		t1, t2 := p1.Next, p2.Next

		p1.Next = p2
		if t1 == nil && t2 != nil {
			break
		}
		p2.Next = t1

		p1 = t1
		p2 = t2
	}
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

//func reverse(head *ListNode) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//
//	tail := reverse(head.Next)
//	//tail.Next = head 返回的tail是最后的节点不是原序当前节点的后继
//	head.Next.Next = nil
//	head.Next = nil
//	return tail
//}
