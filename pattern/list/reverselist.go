package list

// 循环
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	prev := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//
//	return prev
//}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	prevv, currr := dummy, head
	for i := 0; i < left-1; i++ {
		prevv, currr = prevv.Next, currr.Next
	}

	for i := left; i < right && currr.Next != nil; i++ {
		next := currr.Next
		currr.Next = next.Next
		next.Next = prevv.Next
		prevv.Next = next

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
//	prev := dummy
//	curr := head
//	shift := 0
//	for ; shift < left-1; shift++ {
//		prev, curr = prev.Next, curr.Next
//	}
//
//	start := curr
//	for ; shift < right; shift++ {
//		next := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = next
//	}
//
//	start.Next.Next = prev
//	start.Next = curr
//
//	return dummy.Next
//}
