package trees

import (
	"fmt"
	"math"
)

// 递归中序遍历
func isValidBST(root *TreeNode) bool {
	ans := valid(root.Left, math.MinInt, root.Val) && valid(root.Right, root.Val, math.MaxInt)

	return ans
}

func valid(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return valid(root.Left, lower, root.Val) && valid(root.Right, root.Val, upper)
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{}
	latest := math.MinInt

	curr := root
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		fmt.Println(stack)

		curr = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		fmt.Println(stack, curr.Val, latest)

		if curr.Val <= latest {
			return false
		} else {
			latest = curr.Val
		}

		curr = curr.Right
	}

	return true
}
