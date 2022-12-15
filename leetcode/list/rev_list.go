package list

// 反转链表：https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}

func reverseList2(head *ListNode) *ListNode {
	var tail *ListNode

	var helper func(curr *ListNode) *ListNode
	helper = func(curr *ListNode) *ListNode {
		if curr == nil || curr.Next == nil {
			tail = curr
			return curr
		}

		next := helper(curr.Next)
		next.Next = curr
		curr.Next = nil

		return curr
	}

	helper(head)

	return tail
}
