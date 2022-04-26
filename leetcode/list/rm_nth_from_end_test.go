package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	var head *ListNode
	var n int
	var res *ListNode

	t.Run("case1", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2, 3, 4, 5}).Head
		n = 2

		res = removeNthFromEnd(head, n)
		assert.Equal(t, []int{1, 2, 3, 5}, linkedlist.Traverse(res))
	})

	t.Run("case2", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2}).Head
		n = 1

		res = removeNthFromEnd(head, n)
		assert.Equal(t, []int{1}, linkedlist.Traverse(res))
	})

	t.Run("case2", func(t *testing.T) {
		head = linkedlist.FromSlice([]int{1, 2}).Head
		n = 2

		res = removeNthFromEnd(head, n)
		assert.Equal(t, []int{2}, linkedlist.Traverse(res))
	})
}
