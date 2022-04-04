package dynamic

func lengthOfLIS(nums []int) int {
	n := len(nums)
	maxLen := 1

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
