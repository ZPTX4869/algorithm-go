package trees

import (
	"fmt"
	"learn-algorithm/structure/btree"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	vals := []int{61, 46, 66, 43, null, null, null, 39, null, null, null}
	tree := btree.FromSlice(vals)
	fmt.Println(tree.LevelTraverse())

	tree.Root = insertIntoBST(tree.Root, 88)
	fmt.Println(btree.InorderTraverse(tree.Root))
}
