package list

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow, fast = slow.Next, fast.Next
		if fast != nil {
			fast = fast.Next
		}

		if slow == fast {
			break
		}
	}

	if fast != nil {
		for head != slow {
			head, slow = head.Next, slow.Next
		}
		return head
	}

	return nil
}

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//
//	visited := map[*ListNode]struct{}{}
//	for i := 0; head != nil; i++ {
//		if _, ok := visited[head]; ok {
//			return head
//		} else {
//			visited[head] = struct{}{}
//			head = head.Next
//		}
//	}
//
//	return nil
//}
