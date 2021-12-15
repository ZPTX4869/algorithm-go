package binarytree

import (
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	vals := []int{1, 2, 3, null, 4, null, 5}
	tree := FromSlice(vals)

	fmt.Println(tree.Root.Val)
}

func TestBinaryTree_LevelTraverse(t *testing.T) {
	vals := []int{1, 2, 3, null, 4, null, 5}
	tree := FromSlice(vals)

	result := tree.LevelTraverse()
	fmt.Println(result)
}

func TestPreorderTraverse(t *testing.T) {
	vals := []int{1, 2, 3}
	tree := FromSlice(vals)

	result := PreorderTraverse(tree.Root)
	fmt.Println(result)
}

func TestInorderTraverse(t *testing.T) {
	vals := []int{1, 2, 3}
	tree := FromSlice(vals)

	result := InorderTraverse(tree.Root)
	fmt.Println(result)
}

func TestPostorderTraverse(t *testing.T) {
	vals := []int{1, 2, 3}
	tree := FromSlice(vals)

	result := PostorderTraverse(tree.Root)
	fmt.Println(result)
}
