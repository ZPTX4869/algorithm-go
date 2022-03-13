package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var invertTree func(root *TreeNode) *TreeNode
	invertTree = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		invertTree(root.Left)
		invertTree(root.Right)

		root.Left, root.Right = root.Right, root.Left

		return root
	}

	var deepEqual func(r1 *TreeNode, r2 *TreeNode) bool
	deepEqual = func(r1, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		}

		if r1 != nil && r2 != nil {
			if r1.Val != r2.Val {
				return false
			}

			return deepEqual(r1.Left, r2.Left) && deepEqual(r1.Right, r2.Right)
		}

		return false
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	return deepEqual(root.Left, root.Right)
}
