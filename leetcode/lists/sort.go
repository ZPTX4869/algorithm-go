package lists

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	mid := find(head)
	left := mergeSort(head)
	right := mergeSort(mid)
	return merge(left, right)
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

func find(head *ListNode) *ListNode {
	var prev *ListNode
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		prev = slow
		slow = slow.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	return slow
}

//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	var slice []int
//	for curr := head; curr != nil; curr = curr.Next {
//		slice = append(slice, curr.Val)
//	}
//
//	mergeSort(slice, 0, len(slice)-1)
//
//	curr := head
//	for _, val := range slice {
//		curr.Val = val
//		curr = curr.Next
//	}
//
//	return head
//}
//
//func mergeSort(slice []int, left, right int) {
//	if left == right {
//		return
//	}
//
//	// 最好不要写成 (left + right)/2 ，这样可能造成整型溢出
//	mid := left + (right-left)/2
//	mergeSort(slice, left, mid)
//	mergeSort(slice, mid+1, right)
//
//	merge(slice, left, mid, right)
//}
//
//func merge(slice []int, left, mid, right int) {
//	var newSlice []int
//	lp, rp := left, mid+1
//	for lp <= mid && rp <= right {
//		if slice[lp] < slice[rp] {
//			newSlice = append(newSlice, slice[lp])
//			lp++
//		} else {
//			newSlice = append(newSlice, slice[rp])
//			rp++
//		}
//	}
//
//	if lp > mid {
//		newSlice = append(newSlice, slice[rp:right+1]...)
//	} else {
//		newSlice = append(newSlice, slice[lp:mid+1]...)
//	}
//
//	copy(slice[left:right+1], newSlice)
//}
