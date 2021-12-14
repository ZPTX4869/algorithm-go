package btree

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

// PreorderTraverse todo: 完成前序遍历(非递归)
func PreorderTraverse(root *TreeNode) []int {
	return nil
}

func InorderTraverse(root *TreeNode) []int {
	if root == nil {
		panic("Root can not be empty")
	}

	var result []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		cur := stack[len(stack)-1]
		result = append(result, cur.Val)

		if cur.Right != nil {
			root = cur.Right
		}

		stack = stack[0 : len(stack)-1]
	}

	return result
}

// PostorderTraverse todo: 完成后序遍历(非递归)
func PostorderTraverse(root *TreeNode) []int {
	return nil
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
