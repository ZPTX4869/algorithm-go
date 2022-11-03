package array

import "sort"

// Two pointer
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lp, rp := i+1, n-1
		for lp < rp {
			sum := nums[i] + nums[lp] + nums[rp]
			if sum < 0 {
				lp++
			} else if sum > 0 {
				rp--
			} else {
				res = append(res, []int{nums[i], nums[lp], nums[rp]})
				lp, rp = lp+1, rp-1
				for lp < rp && nums[lp] == nums[lp-1] {
					lp++
				}
				for lp < rp && nums[rp] == nums[rp+1] {
					rp--
				}
			}
		}
	}

	return res
}
