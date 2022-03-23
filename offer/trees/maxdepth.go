package trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	getMax := func(x, y int) int{
		if x > y {
			return x
		}

		return y
	}

	return getMax(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}
