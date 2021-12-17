package lists

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := linkedlist.FromSlice([]int{1, 2, 4})
	l2 := linkedlist.FromSlice([]int{1, 3, 4})

	l3 := linkedlist.LinkedList{
		Head: mergeTwoLists(l1.Head, l2.Head),
	}

	fmt.Println(l3.Traverse())
}
