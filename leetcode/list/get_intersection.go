package list

// 链表相交：https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	curA, curB := headA, headB

	for curA != nil {
		curA = curA.Next
		lenA += 1
	}

	for curB != nil {
		curB = curB.Next
		lenB += 1
	}

	if lenA < lenB {
		lenA, lenB = lenB, lenA
		headA, headB = headB, headA
	}
	diff := lenA - lenB

	for i := 0; i < diff; i++ {
		headA = headA.Next
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
