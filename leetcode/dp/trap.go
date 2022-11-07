package dp

import (
	"algorithm-go/util/maths"
)

// 接雨水：https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	lmax, rmax := make([]int, len(height)), make([]int, len(height))

	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = maths.Max(lmax[i-1], height[i])
	}

	rmax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = maths.Max(rmax[i+1], height[i])
	}

	var sum int
	for i, v := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}

		diff := maths.Min(lmax[i], rmax[i]) - v
		if diff > 0 {
			sum += diff
		}
	}

	return sum
}
