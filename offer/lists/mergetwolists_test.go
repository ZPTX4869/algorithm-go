package lists

import (
	"algorithm-go/structure/linkedlist"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := linkedlist.FromSlice([]int{1, 2, 4}).Head
	list2 := linkedlist.FromSlice([]int{1, 3, 4}).Head
	res := mergeTwoLists(list1, list2)

	linkedlist.TraversePrint(res)
}
