package trees

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		res = append(res, curr.Val)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return res
}

func levelOrder2(root *TreeNode) [][]int {
	type nodeInfo struct {
		node  *TreeNode
		level int
	}

	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]nodeInfo, 0)
	queue = append(queue, nodeInfo{
		node:  root,
		level: 0,
	})

	for len(queue) != 0 {
		currNode, currLevel := queue[0].node, queue[0].level
		queue = queue[1:]

		if len(res) < currLevel+1 {
			res = append(res, make([]int, 0))
		}
		res[currLevel] = append(res[currLevel], currNode.Val)

		if currNode.Left != nil {
			queue = append(queue, nodeInfo{
				node:  currNode.Left,
				level: currLevel + 1,
			})
		}

		if currNode.Right != nil {
			queue = append(queue, nodeInfo{
				node:  currNode.Right,
				level: currLevel + 1,
			})
		}
	}

	return res
}

func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		res = append(res, make([]int, 0))
		level := len(res) - 1
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		if len(res)%2 == 0 {
			for i := size - 1; i >= 0; i-- {
				curr := queue[i]
				res[level] = append(res[level], curr.Val)
			}
		} else {
			for i := 0; i < size; i++ {
				curr := queue[i]
				res[level] = append(res[level], curr.Val)
			}
		}

		queue = queue[size:]
	}

	return res
}
