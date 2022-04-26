package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	var h1, h2 *ListNode
	var res *ListNode

	t.Run("case1", func(t *testing.T) {
		h1 = linkedlist.FromSlice([]int{1, 9, 1}).Head
		h2 = linkedlist.FromSlice([]int{3}).Head
		h3 := linkedlist.FromSlice([]int{2, 4}).Head

		h1.Next.Next.Next = h3
		h2.Next = h3

		res = getIntersectionNode(h1, h2)
		assert.Equal(t, h3, res)
	})

	t.Run("case2", func(t *testing.T) {
		h1 = linkedlist.FromSlice([]int{2, 6, 4}).Head
		h2 = linkedlist.FromSlice([]int{1, 5}).Head

		res = getIntersectionNode(h1, h2)
		assert.Equal(t, (*ListNode)(nil), res)
	})
}
