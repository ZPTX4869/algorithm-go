package trees

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	vals := []int{4, 2, 7, 1, 3}
	tree := binarytree.FromSlice(vals)
	fmt.Println(binarytree.LevelTraverse(tree.Root))

	tree.Root = insertIntoBST(tree.Root, 5)
	fmt.Println(binarytree.LevelTraverse(tree.Root))
}

func Test_insertIntoBST2(t *testing.T) {
	vals := []int{4, 2, 7, 1, 3}
	tree := binarytree.FromSlice(vals)
	fmt.Println(binarytree.LevelTraverse(tree.Root))

	tree.Root = insertIntoBST2(tree.Root, 5)
	fmt.Println(binarytree.LevelTraverse(tree.Root))
}
