package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	lnode, rnode := head, head
	for i := 0; i < k; i++ {
		if rnode == nil {
			return lnode
		}
		rnode = rnode.Next
	}
	
	newHead := reversePartial(lnode, rnode)
	lnode.Next = reverseKGroup(rnode, k)

	return newHead
}

func reversePartial(left, right *ListNode) *ListNode {
	var prev, curr *ListNode

	curr = left
	for curr != right {
		next := curr.Next
		curr.Next = prev
		// Update ptrs
		prev = curr
		curr = next
	}

	return prev
}

