package tree

func convertBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	idx := 0
	var update func(node *TreeNode)
	update = func(node *TreeNode) {
		if node == nil {
			return
		}

		update(node.Left)
		node.Val = preSum[len(preSum)-1] - preSum[idx]
		idx += 1
		update(node.Right)
	}
	update(root)

	return root
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)
		root.Val = root.Val + sum
		sum = root.Val
		traverse(root.Left)
	}
	traverse(root)

	return root
}
