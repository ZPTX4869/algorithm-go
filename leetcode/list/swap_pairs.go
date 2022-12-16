package list

import "fmt"

// 两两交换链表中的节点：https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummy.Next = head
	pre, first := dummy, head

	for first != nil && first.Next != nil {
		second := first.Next
		nxt := second.Next

		second.Next = first
		first.Next = nxt
		pre.Next = second

		pre = first
		first = nxt
	}

	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	var helper func(first *ListNode) *ListNode
	helper = func(first *ListNode) *ListNode {
		if first == nil || first.Next == nil {
			return first
		}
		fmt.Printf("[Into] first: %v, second: %v\n", first, first.Next)

		second := first.Next
		first.Next = helper(second.Next)
		second.Next = first
		fmt.Printf("[Swap] %v <-> %v\n", second, first)

		return second
	}

	return helper(head)
}
