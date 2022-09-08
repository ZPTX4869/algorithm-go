package tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	nodes := make([]*TreeNode, 0)
	cur := root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			nodes = append(nodes, cur)
			stack = append(stack, cur)
			cur = cur.Left
		}

		stack = stack[:len(stack)-1]
	}

	prev := nodes[0]
	for _, node := range nodes[1:] {
		prev.Right = node
		prev.Left = nil
		prev = node
	}
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	flatten2(root.Left)
	flatten2(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = right
}
