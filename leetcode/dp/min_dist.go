package dp

import "algorithm-go/util/maths"

// 编辑距离：https://leetcode.cn/problems/edit-distance/submissions/
func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			memo[i][j] = 1 + maths.Min(dp(i, j-1), dp(i-1, j), dp(i-1, j-1))
		}

		return memo[i][j]
	}

	return dp(len(word1)-1, len(word2)-1)
}

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + maths.Min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}
