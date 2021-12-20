package lists

func sortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	lp, rp := left, right
	for lp != nil && rp != nil {
		if lp.Val < rp.Val {
			curr.Next = lp
			lp = lp.Next
		} else {
			curr.Next = rp
			rp = rp.Next
		}
		curr = curr.Next
	}

	if lp == nil {
		curr.Next = rp
	} else {
		curr.Next = lp
	}

	return dummy.Next
}

//递归：自顶向下
func mergeSort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow, fast = slow.Next, fast.Next.Next
	}

	mid := slow
	left := mergeSort(head, mid)
	right := mergeSort(mid, tail)
	return merge(left, right)
}

// 循环：自底向上
//func mergeSort(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	length := 0
//	for cur := head; cur != nil; cur = cur.Next {
//		length++
//	}
//
//	dummy := &ListNode{Next: head}
//	// 外循环不断增加待排序的子链表长度
//	for subLength := 1; subLength < length; subLength <<= 1 {
//		prev, curr := dummy, dummy.Next
//		for curr != nil {
//			p1 := curr
//			for i := 1; i < subLength && curr.Next != nil; i++ {
//				curr = curr.Next
//			}
//
//			p2 := curr.Next
//			curr.Next = nil
//			curr = p2
//			for i := 1; i < subLength && curr != nil && curr.Next != nil; i++ {
//				curr = curr.Next
//			}
//
//			var next *ListNode
//			if curr != nil {
//				next = curr.Next
//				curr.Next = nil
//			}
//
//			prev.Next = merge(p1, p2)
//			for prev.Next != nil {
//				prev = prev.Next
//			}
//			curr = next
//		}
//	}
//	return dummy.Next
//}
