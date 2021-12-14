package lists

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
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
//	pre := dummy
//	cur := dummy.Next
//	for cur != nil && cur.Next != nil {
//		isDup = false
//
//		for cur.Next != nil && cur.Val == cur.Next.Val {
//			isDup = true
//			cur.Next = cur.Next.Next
//		}
//
//		if isDup {
//			pre.Next = cur.Next
//			cur = pre.Next
//		} else {
//			pre = pre.Next
//			cur = pre.Next
//		}
//
//	}
//
//	return dummy.Next
//}
