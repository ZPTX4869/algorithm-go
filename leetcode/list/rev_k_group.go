package list

// K个一组反转链表：https://leetcode.cn/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	left, right := head, head

	for i := 0; i < k; i++ {
		if right == nil {
			return left
		}
		right = right.Next
	}

	newHead := reversePartial(left, right)
	left.Next = reverseKGroup(right, k)

	return newHead
}

// 注意，这里的是一个左闭右开区间：[left, right)
func reversePartial(left, right *ListNode) *ListNode {
	var pre *ListNode
	cur := left

	// right是需要反转的最后一个节点的后继
	for cur != right {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}
