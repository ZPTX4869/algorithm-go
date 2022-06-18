package tree

import (
	"algorithm-go/structure/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	assert.Equal(t, root, binarytree.FromSlice([]int{6, 3, 5, null, 2, 0, null, null, 1}).Root)
}
