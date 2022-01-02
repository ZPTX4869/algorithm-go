package lists

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	l1 := linkedlist.FromSlice([]int{4, 2, 1, 3})
	l2 := linkedlist.LinkedList{
		Head: sortList(l1.Head),
	}

	fmt.Println(l2.Traverse())
}
