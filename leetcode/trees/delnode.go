package trees

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		node := root.Left
		for node.Right != nil {
			node = node.Right
		}
		node.Right = root.Right
		return root.Left
	}

	return root
}
