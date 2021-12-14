package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func FromSlice(vals []int) LinkedList {
	if len(vals) == 0 {
		panic("Vals can not be empty")
	}

	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}

	cur := head
	for _, val := range vals[1:] {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		cur.Next = newNode
		cur = newNode
	}

	return LinkedList{
		Head: head,
	}
}

func (l *LinkedList) Traverse() []int {
	var nums []int

	if l.Head == nil {
		return nums
	}

	cur := l.Head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	return nums
}
