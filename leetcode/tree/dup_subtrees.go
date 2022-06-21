package tree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	memo := make(map[string]int, 0)

	var collect func(root *TreeNode) string
	collect = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		left := collect(root.Left)
		right := collect(root.Right)
		curr := left + "," + right + "," + strconv.Itoa(root.Val)

		if memo[curr] == 1 {
			res = append(res, root)
			memo[curr] += 1
		}

		if memo[curr] == 0 {
			memo[curr] = 1
		}

		return curr
	}

	collect(root)

	return res
}
