package lists

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

func reverseList2(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy
	cur := head
	shift := 0
	for ; shift < left-1; shift++ {
		pre, cur = pre.Next, cur.Next
	}

	start := cur
	for ; shift < right; shift++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	start.Next.Next = pre
	start.Next = cur

	return dummy.Next
}
