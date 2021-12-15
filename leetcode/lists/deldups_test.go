package lists

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 1, 2, 3, 3})
	list.Traverse()

	list.Head = deleteDuplicates(list.Head)
	fmt.Println(list.Traverse())
}

func Test_deleteDuplicates2(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 1, 2, 3, 3})
	list.Traverse()

	list.Head = deleteDuplicates2(list.Head)
	fmt.Println(list.Traverse())
}
