package dynamic

func minPathSum(grid [][]int) int {
	getMin := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	numRows, numCols := len(grid), len(grid[0])
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, numCols)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = getMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[numRows-1][numCols-1]
}
