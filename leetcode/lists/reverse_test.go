package lists

import (
	"fmt"
	"learn-algorithm/structure/linkedlist"
	"testing"
)

func Test_reverseList(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3, 4, 5})
	list.Traverse()

	list.Head = reverseList(list.Head)
	fmt.Println(list.Traverse())
}

func Test_reverseList2(t *testing.T) {
	list := linkedlist.FromSlice([]int{5})
	list.Traverse()

	list.Head = reverseList2(list.Head, 1, 1)
	fmt.Println(list.Traverse())
}
