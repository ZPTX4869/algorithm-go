package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}

	if root == nil {
		return newNode
	}

	curr := root
	for curr != nil {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = newNode
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = newNode
				break
			}
			curr = curr.Right
		}
	}

	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	currr := root
	stack := []*TreeNode{}

	var node *TreeNode
	for len(stack) != 0 || currr != nil {
		for currr != nil {
			stack = append(stack, currr)
			currr = currr.Left
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
			currr = node.Right
		}
	}

	if root == nil {
        root = &TreeNode{Val: val}
	} else {
        node.Right = &TreeNode{Val: val}
    }

	return root
}
