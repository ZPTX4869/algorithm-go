package lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  1<<31 - 1,
		Next: nil,
	}
	
	curr := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}

		curr = curr.Next
	}

	if p1 == nil {
		curr.Next = p2
	} else {
		curr.Next = p1
	}

	return dummy.Next
}
