package lists

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var res *ListNode
	if head == nil {
		return head
	}

	var helper func(curr *ListNode) int
	helper = func(curr *ListNode) int {
		if curr == nil {
			return 0
		}

		cnt := helper(curr.Next) + 1

		if cnt == k {
			res = curr
		}
	
		return cnt
	}

	helper(head)

	return res
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	return slow
}
