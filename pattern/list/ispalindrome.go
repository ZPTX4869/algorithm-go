package list

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	latter := slow.Next
	slow.Next = nil

	var prev *ListNode
	curr := latter
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}

	return true
}

// 将链表转换为数组，然后通过索引从两头往中间遍历
//func isPalindrome(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return true
//	}
//
//	vals := make([]int, 0)
//	for head != nil {
//		vals = append(vals, head.Val)
//		head = head.Next
//	}
//
//	i, j := 0, len(vals)-1
//	for i < j {
//		if vals[i] != vals[j] {
//			return false
//		}
//		i++
//		j--
//	}
//
//	return true
//}

// 找中点的同时将前半截的元素复制到栈中，然后比较每个位置的值是否和链表后半段对应
//func isPalindrome(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return true
//	}
//
//	stack := make([]int, 0)
//	slow, fast := head, head
//	for {
//		if fast.Next == nil {
//			slow = slow.Next
//			break
//		}
//
//		if fast.Next.Next == nil {
//			stack = append(stack, slow.Val)
//			slow = slow.Next
//			break
//		}
//
//		stack = append(stack, slow.Val)
//		slow, fast = slow.Next, fast.Next.Next
//	}
//
//	for len(stack) != 0 {
//		if slow == nil || stack[len(stack)-1] != slow.Val {
//			return false
//		}
//
//		stack = stack[0 : len(stack)-1]
//		slow = slow.Next
//	}
//
//	return true
//}
