package trees

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	vals := []int{61, 46, 66, 43, null, null, null, 39, null, null, null}
	tree := binarytree.FromSlice(vals)
	fmt.Println(tree.LevelTraverse())

	tree.Root = insertIntoBST(tree.Root, 88)
	fmt.Println(tree.InorderTraverse())
}
