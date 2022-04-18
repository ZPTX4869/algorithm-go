package binarytree

import (
	"fmt"
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
		return BinaryTree{
			Root: nil,
		}
	}

	if vals[0] == null {
		panic("root can't be empty")
	}

	var skip int
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}

	i := 1
	for i < len(vals) {
		numNodes := len(queue)

		for j := 0; j < numNodes; j++ {
			for vals[i] == null {
				skip += 1
				i++
				if i == len(vals) {
					goto END
				}
			}

			curr := queue[0]

			newNode := &TreeNode{Val: vals[i]}
			if curr.Left == nil {
				if skip == 0 {
					curr.Left = newNode
					queue = append(queue, curr.Left)
					i++
				} else {
					skip -= 1
				}
			}

			for vals[i] == null {
				skip += 1
				i++
				if i == len(vals) {
					goto END
				}
			}
			newNode = &TreeNode{Val: vals[i]}
			if skip == 0 {
				curr.Right = newNode
				queue = append(queue, curr.Right)
				i++
			} else {
				skip -= 1
			}

			queue = queue[1:]
			if len(queue) == 0 {
				panic(fmt.Sprintf("assign a node { val:%d } to an empty node", vals[j]))
			}
		}
	}
END:
	return BinaryTree{
		Root: root,
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
