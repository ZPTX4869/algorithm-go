package lists

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	largeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	smallHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	prev1, prev2 := largeHead, smallHead
	curr := head
	for curr != nil {
		if curr.Val < x {
			prev1.Next = curr.Next
			prev2.Next = curr

			curr = curr.Next
			prev2 = prev2.Next
		} else {
			prev1, curr = prev1.Next, curr.Next
		}
	}
	prev2.Next = largeHead.Next

	return smallHead.Next
}
