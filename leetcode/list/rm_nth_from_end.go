package list

// 删除链表的倒数第N个结点：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := head, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
