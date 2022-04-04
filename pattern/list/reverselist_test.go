package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3, 4, 5})
	list.Traverse()

	list.Head = reverseList(list.Head)
	fmt.Println(list.Traverse())
}

func Test_reverseList2(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3, 4, 5})
	list.Traverse()

	list.Head = reverseBetween(list.Head, 2, 4)
	fmt.Println(list.Traverse())
}

func Test_reverseList2_ext(t *testing.T) {
	list := linkedlist.FromSlice([]int{3, 5})
	list.Traverse()

	list.Head = reverseBetween(list.Head, 1, 1)
	fmt.Println(list.Traverse())
}
