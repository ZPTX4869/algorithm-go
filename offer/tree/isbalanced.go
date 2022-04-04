package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if rightDepth - leftDepth > 1  || rightDepth - leftDepth < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalanced2(root *TreeNode) bool {
	getMax := func (x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	var check func(root *TreeNode) int
	check = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := check(root.Left)
		if left == -1 {
			return -1
		}

		right := check(root.Right)
		if right == -1 {
			return -1
		}
		
		if left - right < -1 || left - right > 1 {
			return -1
		}

		return getMax(left+1, right+1)
	}

	return check(root) != -1
}