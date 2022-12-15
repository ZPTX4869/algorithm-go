package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	var head *ListNode
	var k int
	var res *ListNode

	t.Run("case1", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2, 3, 4, 5}).Head
		k = 2
		res = reverseKGroup(head, k)
		fmt.Println(linkedlist.Traverse(res))
	})
}

func Test_reversePartial(t *testing.T) {
	var head *ListNode
	var res *ListNode

	t.Run("case1", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2, 3}).Head
		res = reversePartial(head, head.Next.Next.Next)
		fmt.Println(linkedlist.Traverse(res))
	})
}
