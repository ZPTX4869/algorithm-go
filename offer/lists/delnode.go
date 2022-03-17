package lists

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Val: 1 << 31,
	}
	dummy.Next = head

	prev, curr := dummy, head
	for curr != nil {
		if curr.Val == val {
			next := curr.Next
			prev.Next = next
			curr.Next = nil
		}
		prev, curr = curr, curr.Next
	}

	return dummy.Next
}