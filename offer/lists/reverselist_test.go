package lists

import (
	"algorithm-go/structure/linkedlist"
	"testing"
)

func Test_reverseList(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3})
	newHead := reverseList2(list.Head)

	linkedlist.TraversePrint(newHead)
}
