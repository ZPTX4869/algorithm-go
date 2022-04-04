package dynamic

func minimumTotal(triangle [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	numRows := len(triangle)
	dp := make([]int, numRows+1)

	for i := numRows - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

//func minimumTotal(triangle [][]int) int {
//	min := math.MaxInt32
//
//	min = dfs(triangle, 0, 0)
//	//dfs(triangle, 0, 0, 0)
//
//	return min
//}

// 自低向上分治
//func dfs(triangle [][]int, row, col int) int {
//	getMin := func(x, y int) int {
//		if x < y {
//			return x
//		}
//
//		return y
//	}
//
//	if row == len(triangle) {
//		return 0
//	}
//
//	curr := triangle[row][col]
//	return getMin(dfs(triangle, row+1, col)+curr, dfs(triangle, row+1, col+1)+curr)
//}

// 自定向下递归
//func dfs(triangle [][]int, sum, row, col int) {
//	if row == len(triangle) {
//		if sum < min {
//			min = sum
//		}
//		return
//	}
//
//	sum += triangle[row][col]
//	dfs(triangle, sum, row+1, col)
//	dfs(triangle, sum, row+1, col+1)
//}
