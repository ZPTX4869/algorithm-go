package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3, 4, 5})
	reorderList(list.Head)
	fmt.Println(list.Traverse())
}
