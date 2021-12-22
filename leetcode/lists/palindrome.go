package lists

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	stack := make([]int, 0)
	slow, fast := head, head
	for {
		if fast.Next == nil {
			slow = slow.Next
			break
		}

		if fast.Next.Next == nil {
			stack = append(stack, slow.Val)
			slow = slow.Next
			break
		}

		stack = append(stack, slow.Val)
		slow, fast = slow.Next, fast.Next.Next
	}

	for len(stack) != 0 {
		if slow == nil || stack[len(stack)-1] != slow.Val {
			return false
		}

		stack = stack[0 : len(stack)-1]
		slow = slow.Next
	}

	return true
}
