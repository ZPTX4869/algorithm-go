package lists

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode

	var helper func(curr *ListNode) *ListNode
	helper = func(curr *ListNode) *ListNode {
		if curr == nil || curr.Next == nil {
			newHead = curr
			return curr
		}

		next := helper(curr.Next)
		next.Next = curr
		curr.Next = nil

		return curr
	}

	helper(head)

	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, curr *ListNode
	curr = head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

