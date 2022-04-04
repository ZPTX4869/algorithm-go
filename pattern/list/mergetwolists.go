package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	curr := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}

		curr = curr.Next
	}

	if p1 == nil {
		curr.Next = p2
	} else {
		curr.Next = p1
	}

	return dummy.Next
}

// 递归求解
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//
//	if list2 == nil {
//		return list1
//	}
//
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}
