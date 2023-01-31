package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	tree := FromSlice([]int{1, 2, Null, Null, 3, Null, Null})
	ans := LevelTraverse(tree.Root)

	assert.Equal(t, 3, tree.Size())
	assert.Equal(t, []int{1, 2, 3}, ans)
}

func TestLevelTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, Null, 4, Null, 5})
	ans := tree.LevelTraverse()

	assert.Equal(t, []int{1, 2, 3, 4, 5}, ans)
}

func TestPreorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, Null, 4, Null, 5})
	ans := tree.PreorderTraverse()

	assert.Equal(t, []int{1, 2, 4, 3, 5}, ans)
}

func TestInorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, Null, 4, Null, 5})
	ans := tree.InorderTraverse()

	assert.Equal(t, []int{2, 4, 1, 3, 5}, ans)
}

func TestPostorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, Null, 4, Null, 5})
	ans := tree.PostorderTraverse()

	assert.Equal(t, []int{4, 2, 5, 3, 1}, ans)
}
