package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	var l1, l2 linkedlist.LinkedList
	var res *linkedlist.ListNode

	t.Run("case1", func(t *testing.T) {
		l1 = linkedlist.FromSlice([]int{1, 2, 4})
		l2 = linkedlist.FromSlice([]int{1, 3, 4})

		res = mergeTwoLists(l1.Head, l2.Head)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, linkedlist.Traverse(res))
	})
}
