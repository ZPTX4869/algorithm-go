package tree

import (
	"algorithm-go/structure/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree := binarytree.FromSlice([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4})
	p := binarytree.LevelSearch(tree.Root, 5)
	q := binarytree.LevelSearch(tree.Root, 1)

	ancestor := lowestCommonAncestor(tree.Root, p, q)
	assert.Equal(t, 3, ancestor.Val)
}
