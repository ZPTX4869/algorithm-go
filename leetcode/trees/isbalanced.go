package trees

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := deep(root.Left) - deep(root.Right)
	if math.Abs(float64(diff)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(deep(root.Left), deep(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	
	return y
}
