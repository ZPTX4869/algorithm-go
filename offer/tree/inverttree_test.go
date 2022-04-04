package tree

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_invertTree(t *testing.T) {
	root := binarytree.FromSlice([]int{4, 2, 7, 1, 3, 6, 9}).Root
	res := binarytree.LevelTraverse(root)
	fmt.Println(res)

	newRoot := invertTree(root)
	res = binarytree.LevelTraverse(newRoot)
	fmt.Println(res)
}
