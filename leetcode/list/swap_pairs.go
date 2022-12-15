package list

// 两两交换链表中的节点：https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummy.Next = head
	pre, first := dummy, head

	for first != nil && first.Next != nil {
		second := first.Next
		nxt := second.Next

		second.Next = first
		first.Next = nxt
		pre.Next = second

		pre = first
		first = nxt
	}

	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummy.Next = head
	pre, first := dummy, head
	for first != nil && first.Next != nil {
		second := first.Next
		nxt := second.Next

		second.Next = first
		first.Next = nxt
		pre.Next = second

		pre = first
		first = nxt
	}

	return dummy.Next
}
