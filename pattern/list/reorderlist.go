package list

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//var prev *ListNode
	//slow, fast := head, head
	//for fast != nil && fast.Next != nil {
	//	prev = slow
	//	slow, fast = slow.Next, fast.Next.Next
	//}
	//prev.Next = nil
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	next := slow.Next
	slow.Next = nil
	tail := reverse(next)

	p1, p2 := head, tail
	for p1 != nil {
		t1, t2 := p1.Next, p2.Next

		p1.Next = p2
		p2.Next = t1

		if t2 == nil {
			break
		}

		p1 = t1
		p2 = t2
	}
}

//循环反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 递归反转链表
//func reverse(head *ListNode) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//
//	tail := reverse(head.Next)
//	//tail.Next = head 返回的tail是最后的节点不是原序当前节点的后继
//	head.Next.Next = nil
//	head.Next = nil
//	return tail
//}

// 利用切片求解
//func reorderList(head *ListNode) {
//	if head == nil || head.Next == nil {
//		return
//	}
//
//	var nodes []*ListNode
//	for curr := head; curr != nil; curr = curr.Next {
//		nodes = append(nodes, curr)
//	}
//
//	i, j := 0, len(nodes)-1
//	for i < j {
//		nodes[i].Next = nodes[j]
//		i++
//		if i == j {
//			break
//		}
//		nodes[j].Next = nodes[i]
//		j--
//	}
//	nodes[i].Next = nil
//
//	return
//}
