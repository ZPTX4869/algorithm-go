package dp

import "algorithm-go/util/maths"

// 让字符串成为回文串的最小插入次数：https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func minInsertions(s string) int {
	n := len(s)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i >= j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s[i] == s[j] {
			memo[i][j] = dp(i+1, j-1)
		} else {
			memo[i][j] = 1 + maths.Min(dp(i+1, j), dp(i, j-1))
		}

		return memo[i][j]
	}

	return dp(0, len(s)-1)
}

func minInsertions2(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 0
	}

	for i := n-1; i >= 0; i-- {
		for j := i+1 ; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = 1 + maths.Min(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
