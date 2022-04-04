package tree

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

		if r1 == nil || r2 == nil {
			return false
		}

		return r1.Val == r2.Val && deepEqual(r1.Left, r2.Left) && deepEqual(r1.Right, r2.Right)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	return deepEqual(root.Left, root.Right)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var deepEqual func(r1 *TreeNode, r2 *TreeNode) bool
	deepEqual = func(r1, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		}

		if r1 == nil || r2 == nil {
			return false
		}

		return r1.Val == r2.Val && deepEqual(r1.Left, r2.Right) && deepEqual(r1.Right, r2.Left)
	}

	return deepEqual(root, root)
}

func isSymmetric3(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)

	for len(queue) >= 2 {
		n1, n2 := queue[0], queue[1]
		queue = queue[2:]

		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)
		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
	}

	return true
}
