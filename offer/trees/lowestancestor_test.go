package trees

import (
	"algorithm-go/structure/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree := binarytree.FromSlice([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4})
	res := lowestCommonAncestor2(tree.Root, tree.Root.Left, tree.Root.Right)

	assert.Equal(t, 3, res.Val)
}
