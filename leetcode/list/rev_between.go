package list

var successor *ListNode

// 反转链表II：https://leetcode.cn/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseTopK(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)

	return head
}

func reverseTopK(head *ListNode, k int) *ListNode {
	if k == 1 {
		successor = head.Next
		return head
	}

	tail := reverseTopK(head.Next, k-1)
	head.Next.Next = head
	head.Next = successor

	return tail
}
