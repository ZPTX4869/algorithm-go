package tree

import "math"

func isValidBST(root *TreeNode) bool {
	var check func(curr *TreeNode, lower, upper int) bool
	check = func(curr *TreeNode, lower, upper int) bool {
		if curr == nil {
			return true
		}

		if curr.Val <= lower || curr.Val >= upper {
			return false
		}

		return check(curr.Left, lower, curr.Val) && check(curr.Right, curr.Val, upper)
	}

	return check(root, -math.MaxInt, math.MaxInt)
}
