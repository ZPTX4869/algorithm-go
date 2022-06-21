package tree

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	tree := binarytree.FromSlice([]int{1, 2, 3, 4, null, 2, 4, null, null, 4})
	result := findDuplicateSubtrees(tree.Root)
	for i := 0; i < len(result); i++ {
		fmt.Println(binarytree.LevelTraverse(result[i]))
	}
}
