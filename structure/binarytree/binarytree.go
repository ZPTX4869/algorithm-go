package binarytree

import (
	"algorithm-go/util/slices"
	"math"
)

const Null = math.MaxInt32

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

func (t *BinaryTree) LevelTraverse() []int {
	return LevelTraverse(t.Root)
}

func (t *BinaryTree) PreorderTraverse() []int {
	return PreorderTraverse(t.Root)
}

func (t *BinaryTree) InorderTraverse() []int {
	return InorderTraverse(t.Root)
}

func (t *BinaryTree) PostorderTraverse() []int {
	return PostorderTraverse(t.Root)
}

func FromSlice(vals []int) BinaryTree {
	if len(vals) == 0 {
		return BinaryTree{
			Root: nil,
		}
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}

	i := 1
	skip := 0

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
		size: len(slices.Filter(vals, func(x int) bool { return x != Null })),
	}
}

const (
	White = iota
	Black
)

type pair struct {
	node  *TreeNode
	color int
}

func PreorderTraverse(root *TreeNode) []int {
	ans := make([]int, 0)
	stk := make([]*pair, 0)

	stk = append(stk, &pair{root, White})

	for len(stk) > 0 {
		cur := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if cur.color == Black {
			ans = append(ans, cur.node.Val)
		} else {
			if cur.node.Right != nil {
				stk = append(stk, &pair{cur.node.Right, White})
			}
			if cur.node.Left != nil {
				stk = append(stk, &pair{cur.node.Left, White})
			}
			stk = append(stk, &pair{cur.node, Black})
		}
	}

	return ans
}

func InorderTraverse(root *TreeNode) []int {
	ans := make([]int, 0)
	stk := make([]*pair, 0)

	stk = append(stk, &pair{root, White})

	for len(stk) > 0 {
		cur := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if cur.color == Black {
			ans = append(ans, cur.node.Val)
		} else {
			if cur.node.Right != nil {
				stk = append(stk, &pair{cur.node.Right, White})
			}
			stk = append(stk, &pair{cur.node, Black})
			if cur.node.Left != nil {
				stk = append(stk, &pair{cur.node.Left, White})
			}
		}
	}

	return ans
}

func PostorderTraverse(root *TreeNode) []int {
	ans := make([]int, 0)
	stk := make([]*pair, 0)

	stk = append(stk, &pair{root, White})

	for len(stk) > 0 {
		cur := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if cur.color == Black {
			ans = append(ans, cur.node.Val)
		} else {
			stk = append(stk, &pair{cur.node, Black})
			if cur.node.Right != nil {
				stk = append(stk, &pair{cur.node.Right, White})
			}
			if cur.node.Left != nil {
				stk = append(stk, &pair{cur.node.Left, White})
			}
		}
	}

	return ans
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

func LevelSearch(root *TreeNode, tar int) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		len := len(queue)

		for i := 0; i < len; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.Val == tar {
				return curr
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

	}

	return nil
}
