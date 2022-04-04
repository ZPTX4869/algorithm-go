package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool, 0)

	curr := headA
	for curr != nil {
		visited[curr] = true
		curr = curr.Next
	}

	curr = headB
	for curr != nil {
		if _, ok := visited[curr]; ok {
			return curr
		}
		curr = curr.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != nil || pb != nil {
		if pa == pb {
			return pa
		}

		if pa == nil {
			pa = headB
			continue
		}
		if pb == nil {
			pb = headA
			continue
		}

		pa = pa.Next
		pb = pb.Next
	}

	return nil
}
