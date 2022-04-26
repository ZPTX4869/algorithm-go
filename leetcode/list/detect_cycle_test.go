package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	var head *ListNode
	var res *ListNode

	t.Run("case1", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{3, 2, 0, -4}).Head
		head.Next.Next.Next.Next = head.Next

		res = detectCycle(head)
		assert.Equal(t, head.Next, res)
	})

	t.Run("case2", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2}).Head
		head.Next.Next = head

		res = detectCycle(head)
		assert.Equal(t, head, res)
	})

	t.Run("case3", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1}).Head

		res = detectCycle(head)
		assert.Equal(t, (*ListNode)(nil), res)
	})

	t.Run("case4", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{}).Head

		res = detectCycle(head)
		assert.Equal(t, (*ListNode)(nil), res)
	})
}
