package dynamics

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}