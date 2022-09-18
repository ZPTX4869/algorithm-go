package dp

import (
	"algorithm-go/util/maths"
)

// 打家劫舍III（树形DP）：https://leetcode.cn/problems/house-robber-iii/
func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	var dp func(node *TreeNode) int

	dp = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if v, ok := memo[node]; ok {
			return v
		}

		case1 := dp(node.Left) + dp(node.Right)
		case2 := node.Val

		if node.Left != nil {
			case2 += dp(node.Left.Left) + dp(node.Left.Right)
		}
		if node.Right != nil {
			case2 += dp(node.Right.Left) + dp(node.Right.Right)
		}

		memo[node] = maths.Max(case1, case2)

		return memo[node]
	}

	return dp(root)
}
