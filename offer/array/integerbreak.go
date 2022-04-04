package array

func integerBreak(n int) int {
	getMax := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0

	for i := 2; i <= n; i++ {
		currMax := 0
		for j := 1; j < i; j++ {
			temp := getMax((i-j)*dp[j], (i-j)*j)
			currMax = getMax(temp, currMax)
		}
		dp[i] = currMax
	}

	return dp[n]
}
