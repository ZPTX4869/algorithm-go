package tree

func kthLargest(root *TreeNode, k int) int {
	res := make([]int, k+1)
	
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		res[k] = curr.Val
		for i := k-1; i >= 0; i-- {
			if res[i] < res[i+1] {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return res[k-1]
}

func kthLargest2(root *TreeNode, k int) int {
	res := make([]int, 0)

	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraverse(root.Left)
		res = append(res, root.Val)
		inorderTraverse(root.Right)
	}

	inorderTraverse(root)

	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res[k-1]
}