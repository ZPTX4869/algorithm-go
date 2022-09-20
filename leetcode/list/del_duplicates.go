package list

// 删除链表中的重复元素II：https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head

	for cur != nil {
		// 如果出现重复，那么移动到该重复节点的最后一个
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			// 如果pre的下一个节点没变，说明没有重复元素，pre后移一个节点
			pre = pre.Next
		} else {
			// 如果pre的下一个节点变了，说明有重复元素，pre的后一个节点指向重复节点的下一个
			pre.Next = cur.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	var delete func(head *ListNode) *ListNode
	delete = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		// 如果出现重复，一直往后找，直到发现第一个不为重复值的节点，返回改节点的递归结果
		if head.Val == head.Next.Val {
			nxt := head.Next
			for nxt != nil && nxt.Val == head.Val {
				nxt = nxt.Next
			}
			return delete(nxt)
		}

		head.Next = delete(head.Next)

		return head
	}

	return delete(head)
}
