package trees

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}

	if root == nil {
		return newNode
	}

	cur := root
	for cur != nil {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = newNode
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = newNode
				break
			}
			cur = cur.Right
		}
	}

	return root
}
