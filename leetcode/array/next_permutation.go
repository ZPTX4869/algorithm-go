package array

import "sort"

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
}
