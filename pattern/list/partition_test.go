package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	l1 := linkedlist.FromSlice([]int{1, 4, 3, 2, 5, 2})
	l2 := linkedlist.LinkedList{
		Head: partition(l1.Head, 3),
	}

	fmt.Println(l2.Traverse())
}
