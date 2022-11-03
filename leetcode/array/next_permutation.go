package array

import "sort"

// 下一个排列：https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}

			sort.Ints(nums[i:])
			return
		}
	}

	// Reverse the whole slice if next permutation doesn't exist.
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
