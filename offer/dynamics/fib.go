package dynamics

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-2] + dp[i-1]
		dp[i] = dp[i] % 1000000007
	}

	return dp[n]
}
