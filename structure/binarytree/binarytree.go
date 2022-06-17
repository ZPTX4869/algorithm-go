package binarytree

import (
	"algorithm-go/util"
	"math"
)

const Null = math.MaxInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
	size int
}

func (t BinaryTree) Size() int {
	return t.size
}

func FromSlice(vals []int) BinaryTree {
	if len(vals) == 0 {
		return BinaryTree{
			Root: nil,
		}
	}

	var skip int
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}

	i := 1
	for i < len(vals) {
		if vals[i] == Null {
			skip += 1
			i++
			continue
		}

		curr := queue[0]
		newNode := &TreeNode{Val: vals[i]}

		if curr.Left == nil {
			if skip == 0 {
				curr.Left = newNode
				queue = append(queue, curr.Left)
				i++
				continue
			} else {
				skip -= 1
			}
		}

		if skip == 0 {
			curr.Right = newNode
			queue = append(queue, curr.Right)
			i++
		} else {
			skip -= 1
		}

		queue = queue[1:]
	}

	return BinaryTree{
		Root: root,
		size: len(util.FilterSlice(vals, func(x int) bool { return x != Null })),
	}
}

func LevelTraverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		len := len(queue)

		for i := 0; i < len; i++ {
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
