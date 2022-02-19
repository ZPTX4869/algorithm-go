package dynamics

import "fmt"

func backPack(m int, A []int) int {
	n := len(A)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
		}
		fmt.Printf("%d: %v\n", i, dp[i][1:])
	}

	res := 0
	for i := m; i >= 0; i-- {
		if dp[n][i] {
			res = i
			break
		}
	}

	return res
}
