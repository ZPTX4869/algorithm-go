package dynamics

func canJump(nums []int) bool {
	size := len(nums)
	dp := make([]bool, size)
	dp[0] = true

	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			dist := i - j
			if dp[j] && nums[j] >= dist {
				dp[i] = true
				break
			}
		}
	}

	return dp[size-1]
}

func jump(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[0] = 0

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			dist := i - j
			if nums[j] >= dist {
				dp[i] = dp[j] + 1
				break
			}
		}
	}

	return dp[size-1]
}
