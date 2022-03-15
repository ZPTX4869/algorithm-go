package dynamics

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	getMax := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = getMax(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}

	return dp[m][n]
}

func maxValue2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	getMax := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	m, n := len(grid), len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += getMax(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
