package trees

import (
	"algorithm-go/structure/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	res := binarytree.LevelTraverse(root)

	assert.Equal(t, []int{3, 9, 20, 15, 7}, res)
}
