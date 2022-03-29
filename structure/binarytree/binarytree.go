package binarytree

import (
	"math"
)

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
		curr := queue[0]

		if curr.Left == nil {
			curr.Left = newNode
		} else if curr.Right == nil {
			curr.Right = newNode
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
		curr := queue[0]

		if curr.Left != nil {
			if curr.Left.Val == null {
				curr.Left = nil
			} else {
				queue = append(queue, curr.Left)
			}
		}
		if curr.Right != nil {
			if curr.Right.Val == null {
				curr.Right = nil
			} else {
				queue = append(queue, curr.Right)
			}
		}

		queue = queue[1:]
	}
}

func LevelTraverse(root *TreeNode) []int {
	if root == nil {
		panic("Root can not be nil!")
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr.Val)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return result
}

func PreorderTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) != 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		currr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currr.Right != nil {
			root = currr.Right
		}
	}

	return res
}

func InorderTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		currr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, currr.Val)

		if currr.Right != nil {
			root = currr.Right
		}
	}

	return res
}

func PostorderTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisited *TreeNode

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		curr := stack[len(stack)-1]

		if curr.Right != nil && curr.Right != lastVisited {
			root = curr.Right
		} else {
			stack = stack[:len(stack)-1]
			lastVisited = curr
			res = append(res, curr.Val)
		}
	}

	return res
}
