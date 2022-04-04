package list

func reversePrint(head *ListNode) []int {
	var res []int
	
	var helper func (head *ListNode)
	helper = func(head *ListNode) {
		if head == nil {
			return
		}

		helper(head.Next)
		res = append(res, head.Val)
	}

	helper(head)

	return res
}

func reversePrint2(head *ListNode) []int {
	var res []int
	var stk []int

	for head != nil {
		stk = append(stk, head.Val)
		head = head.Next
	}

	for i := len(stk)-1; i >= 0; i-- {
		res = append(res, stk[i])
	}
	
	return res
}