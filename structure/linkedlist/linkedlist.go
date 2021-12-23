package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func FromSlice(vals []int) LinkedList {
	if len(vals) == 0 {
		panic("Vals can not be empty")
	}

	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}

	curr := head
	for _, val := range vals[1:] {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		curr.Next = newNode
		curr = newNode
	}

	return LinkedList{
		Head: head,
		Tail: curr,
	}
}

func (l *LinkedList) Traverse() []int {
	var nums []int

	curr := l.Head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	return nums
}

func TraversePrint(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}

func ReversePrint(head *ListNode) {
	if head.Next == nil {
		fmt.Printf("%v ", head.Val)
		return
	}
	ReversePrint(head.Next)
	fmt.Printf("%v ", head.Val)
}
