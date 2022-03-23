package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, null, 4, null, 5})
	ans := LevelTraverse(tree.Root)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, ans)
}

func TestLevelTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, null, 4, null, 5})
	ans := LevelTraverse(tree.Root)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, ans)
}

func TestPreorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, null, 4, null, 5})
	ans := PreorderTraverse(tree.Root)

	assert.Equal(t, []int{1, 2, 4, 3, 5}, ans)
}

func TestInorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, null, 4, null, 5})
	ans := InorderTraverse(tree.Root)

	assert.Equal(t, []int{2, 4, 1, 3, 5}, ans)
}

func TestPostorderTraverse(t *testing.T) {
	tree := FromSlice([]int{1, 2, 3, null, 4, null, 5})
	ans := PostorderTraverse(tree.Root)

	assert.Equal(t, []int{4, 2, 5, 3, 1}, ans)
}
