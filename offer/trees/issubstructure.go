package trees

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, A)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Val == B.Val {
			if deepEqual(curr.Left, B.Left) && deepEqual(curr.Right, B.Right) {
				return true
			}
		}

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return false
}

func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return deepEqual(A, B) || isSubStructure2(A.Left, B) || isSubStructure2(A.Right, B)
}

func deepEqual(curr *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if curr == nil {
		return false
	}

	if curr.Val != B.Val {
		return false
	}

	return deepEqual(curr.Left, B.Left) && deepEqual(curr.Right, B.Right)

}
