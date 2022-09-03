package list

import (
	"container/heap"
)

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(node interface{}) {
	if node.(*ListNode) == nil {
		return
	}

	*h = append(*h, node.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	n := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]

	return n
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	minHeap := &MinHeap{}
	for _, v := range lists {
		heap.Push(minHeap, v)
	}

	curr := dummy
	for minHeap.Len() != 0 {
		if node := heap.Pop(minHeap).(*ListNode); node != nil {
			if node.Next != nil {
				heap.Push(minHeap, node.Next)
			}

			curr.Next = node
			curr = curr.Next
		}
	}
	return dummy.Next
}
