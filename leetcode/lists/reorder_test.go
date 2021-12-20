package lists

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	tree := linkedlist.FromSlice([]int{1, 2, 3, 4, 5})
	reorderList(tree.Head)
	fmt.Println(tree.Traverse())
}
