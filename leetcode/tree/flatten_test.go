package tree

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	root := binarytree.FromSlice([]int{1, 2, 5, 3, 4, null, 6}).Root
	flatten(root)
	res := binarytree.LevelTraverse(root)
	fmt.Println(res)
}

func Test_flatten2(t *testing.T) {
	root := binarytree.FromSlice([]int{1, 2, 5, 3, 4, null, 6}).Root
	flatten2(root)
	res := binarytree.LevelTraverse(root)
	fmt.Println(res)
}
