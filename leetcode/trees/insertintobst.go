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

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	curr := root
	stack := []*TreeNode{}

	var node *TreeNode
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		node = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if node.Val > val {
			newNode := &TreeNode{Val: val}
			tmp := node.Left
			node.Left = newNode
			newNode.Left = tmp

			return root
		}

		if node.Right != nil {
			curr = node.Right
		}
	}

	if root == nil {
        root = &TreeNode{Val: val}
	} else {
        node.Right = &TreeNode{Val: val}
    }

	return root
}
