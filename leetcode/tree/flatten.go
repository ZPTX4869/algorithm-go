package tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	nodes := make([]*TreeNode, 0)
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			nodes = append(nodes, curr)
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Right
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
