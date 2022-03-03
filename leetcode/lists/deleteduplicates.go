package lists

// 保留一个重复的节点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

// 删除所有重复的节点
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

//func deleteDuplicates2(head *ListNode) *ListNode {
//	dummy := &ListNode{
//		Val:  math.MaxInt64,
//		Next: head,
//	}
//
//	var isDup bool
//	prev := dummy
//	curr := dummy.Next
//	for curr != nil && curr.Next != nil {
//		isDup = false
//
//		for curr.Next != nil && curr.Val == curr.Next.Val {
//			isDup = true
//			curr.Next = curr.Next.Next
//		}
//
//		if isDup {
//			prev.Next = curr.Next
//			curr = prev.Next
//		} else {
//			prev = prev.Next
//			curr = prev.Next
//		}
//
//	}
//
//	return dummy.Next
//}
