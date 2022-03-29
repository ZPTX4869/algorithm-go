package lists

func sortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	currr := dummy

	lp, rp := left, right
	for lp != nil && rp != nil {
		if lp.Val < rp.Val {
			currr.Next = lp
			lp = lp.Next
		} else {
			currr.Next = rp
			rp = rp.Next
		}
		currr = currr.Next
	}

	if lp == nil {
		currr.Next = rp
	} else {
		currr.Next = lp
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
//	for curr := head; curr != nil; curr = curr.Next {
//		length++
//	}
//
//	dummy := &ListNode{Next: head}
//	// 外循环不断增加待排序的子链表长度
//	for subLength := 1; subLength < length; subLength <<= 1 {
//		prev, currr := dummy, dummy.Next
//		for currr != nil {
//			p1 := currr
//			for i := 1; i < subLength && currr.Next != nil; i++ {
//				currr = currr.Next
//			}
//
//			p2 := currr.Next
//			currr.Next = nil
//			currr = p2
//			for i := 1; i < subLength && currr != nil && currr.Next != nil; i++ {
//				currr = currr.Next
//			}
//
//			var next *ListNode
//			if currr != nil {
//				next = currr.Next
//				currr.Next = nil
//			}
//
//			prev.Next = merge(p1, p2)
//			for prev.Next != nil {
//				prev = prev.Next
//			}
//			currr = next
//		}
//	}
//	return dummy.Next
//}
