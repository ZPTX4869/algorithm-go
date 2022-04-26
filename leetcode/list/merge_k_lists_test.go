package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	var lists []*ListNode

	t.Run("case1", func(t *testing.T) {
		list1 := linkedlist.FromSlice([]int{1, 4, 5})
		list2 := linkedlist.FromSlice([]int{1, 3, 4})
		list3 := linkedlist.FromSlice([]int{2, 6})
		lists = []*ListNode{list1.Head, list2.Head, list3.Head}

		res := mergeKLists(lists)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, linkedlist.Traverse(res))
	})

	t.Run("case2", func(t *testing.T) {
		lists = []*ListNode{nil}

		res := mergeKLists(lists)
		assert.Equal(t, []int{}, linkedlist.Traverse(res))
	})
}
