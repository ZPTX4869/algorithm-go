package binarytree

import "math"

const null = math.MaxInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func FromSlice(vals []int) BinaryTree {
	if len(vals) == 0 {
		panic("Value slice for initialization can not be empty!")
	}

	root := &TreeNode{
		Val: vals[0],
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for _, val := range vals[1:] {
		newNode := &TreeNode{Val: val}
		cur := queue[0]

		if cur.Left == nil {
			cur.Left = newNode
		} else if cur.Right == nil {
			cur.Right = newNode
			queue = queue[1:]
		}

		// We can not append `null` into queue, it's just for holding place.
		if newNode.Val != null {
			queue = append(queue, newNode)
		}
	}

	// Create by level have to  use `null` as placeholder, and `null` node should be delelted before return.
	delNull(root)

	return BinaryTree{
		Root: root,
	}
}

func delNull(root *TreeNode) {
	if root == nil {
		panic("Root can not be nil!")
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		cur := queue[0]

		if cur.Left != nil {
			if cur.Left.Val == null {
				cur.Left = nil
			} else {
				queue = append(queue, cur.Left)
			}
		}
		if cur.Right != nil {
			if cur.Right.Val == null {
				cur.Right = nil
			} else {
				queue = append(queue, cur.Right)
			}
		}

		queue = queue[1:]
	}
}

func (t *BinaryTree) LevelTraverse() []int {
	var result []int

	if t.Root == nil {
		panic("Root can not be nil!")
	}

	var queue []*TreeNode
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		cur := queue[0]
		result = append(result, cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		queue = queue[1:]
	}

	return result
}

func (t *BinaryTree) PreorderTraverse() []int {
	return nil
}

func (t *BinaryTree) InorderTraverse() []int {
	var result []int

	root := t.Root
	if root == nil {
		return result
	}

	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		curr := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		result = append(result, curr.Val)
		root = curr.Right
	}

	return result
}

func (t *BinaryTree) PostorderTraverse() []int {
	return nil
}

func LevelTraverse(root *TreeNode) []int {
	if root == nil {
		panic("Root can not be nil!")
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return result
}
