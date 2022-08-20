package dp

import (
	"algorithm-go/util/maths"
	"math"
)

// 最大子数组：https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = maths.Max(nums[i], dp[i-1]+nums[i])
	}

	return maths.Max(dp...)
}

func maxSubArray2(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	res := math.MinInt
	minVal := math.MaxInt
	for i := 0; i < n; i++ {
		minVal = maths.Min(minVal, preSum[i])
		res = maths.Max(res, preSum[i+1]-minVal)
	}

	return res
}
