package trees

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	if root == nil {
		return res
	}

	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		if root == q || root == p {
			return true
		}

		return dfs(root.Left) || dfs(root.Right)
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == p || curr == q {
			if dfs(curr.Left) || dfs(curr.Right) {
				return curr
			}

			return res
		}

		if dfs(curr.Left) && dfs(curr.Right) {
			res = curr
		}

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return res
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	if root == nil {
		return res
	}

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root == p || root == q {
			return root
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left != nil && right != nil {
			return root
		} else if left != nil && right == nil {
			return left
		} else if left == nil && right != nil {
			return right
		}

		return nil
	}

	return dfs(root)
}
