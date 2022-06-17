package arr

// Spiral Matrix II: https://leetcode.cn/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1

	num := 1
	i, j := 0, 0
	for num <= n*n {
		for j <= right {
			res[i][j] = num
			num += 1
			j++
		}
		up += 1
		j--
		i++

		for i <= down {
			res[i][j] = num
			num += 1
			i++
		}
		right -= 1
		i--
		j--

		for j >= left {
			res[i][j] = num
			num += 1
			j--
		}
		down -= 1
		j++
		i--

		for i >= up {
			res[i][j] = num
			num += 1
			i--
		}
		left += 1
		i++
		j++
	}

	return res
}
