package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	var head *ListNode
	var res bool

	t.Run("case1", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{3, 2, 0, -4}).Head
		head.Next.Next.Next.Next = head.Next

		res = hasCycle(head)
		assert.True(t, res)
	})

	t.Run("case2", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2}).Head
		head.Next.Next = head

		res = hasCycle(head)
		assert.True(t, res)
	})
}
