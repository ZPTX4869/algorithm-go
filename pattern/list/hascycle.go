package list

// 快慢指针法
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil {
		slow, fast = slow.Next, fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if slow == fast {
			return true
		}
	}

	return false
}

// 哈希表法
//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//
//	visited := map[*ListNode]struct{}{}
//	for head != nil {
//		if _, ok := visited[head]; ok {
//			return true
//		}
//
//		visited[head] = struct{}{}
//		head = head.Next
//	}
//
//	return false
//}
