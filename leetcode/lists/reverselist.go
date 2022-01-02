package lists

// 循环
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	pre := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//
//	return pre
//}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	prev, curr := dummy, head
	for i := 0; i < left-1; i++ {
		prev, curr = prev.Next, curr.Next
	}

	for i := left; i < right && curr.Next != nil; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next

		if next.Val == right {
			break
		}
	}

	return dummy.Next
}

// 逐个反转法
//func reverseBetween(head *ListNode, left int, right int) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//
//	dummy := &ListNode{
//		Val:  0,
//		Next: head,
//	}
//
//	pre := dummy
//	cur := head
//	shift := 0
//	for ; shift < left-1; shift++ {
//		pre, cur = pre.Next, cur.Next
//	}
//
//	start := cur
//	for ; shift < right; shift++ {
//		next := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//
//	start.Next.Next = pre
//	start.Next = cur
//
//	return dummy.Next
//}
