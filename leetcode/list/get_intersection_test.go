package list

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	list1 := linkedlist.FromSlice([]int{4, 1})
	list2 := linkedlist.FromSlice([]int{5, 0, 1})

	sharedList := linkedlist.FromSlice([]int{8, 4, 5})
	list1.Tail.Next = sharedList.Head
	list2.Tail.Next = sharedList.Head

	ans := getIntersectionNode(list1.Head, list2.Head)
	assert.Equal(t, sharedList.Head, ans)
}
