package dynamics

import "math"

var min int

func init() {
	min = math.MaxInt32
}

func minimumTotal(triangle [][]int) int {
	min = dfs(triangle, 0, 0)
	//dfs(triangle, 0, 0, 0)

	return min
}

// 自低向上分治
func dfs(triangle [][]int, row, col int) int {
	getMin := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	if row == len(triangle) {
		return 0
	}

	cur := triangle[row][col]
	return getMin(dfs(triangle, row+1, col)+cur, dfs(triangle, row+1, col+1)+cur)
}

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
