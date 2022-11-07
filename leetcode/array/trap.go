package array

import "algorithm-go/util/maths"

// 接雨水：https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	var sum int

	for i, v := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}

		var lmax, rmax int
		for _, left := range height[:i] {
			if left > lmax {
				lmax = left
			}
		}
		for _, right := range height[i+1:] {
			if right > rmax {
				rmax = right
			}
		}

		diff := maths.Min(lmax, rmax) - v
		if diff > 0 {
			sum += diff
		}
	}

	return sum
}
