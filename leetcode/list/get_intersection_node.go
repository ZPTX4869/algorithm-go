package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	r1, r2 := true, true

	for p1 !=nil && p2 != nil {
		if p1 == p2 {
			return p1
		}

		p1 = p1.Next
		p2 = p2.Next

		if p1 == nil && r1 {
			p1 = headB
			r1 = false
		}
		if p2 == nil && r2 {
			p2 = headA
			r2 = false
		}
	}

	return nil
}
