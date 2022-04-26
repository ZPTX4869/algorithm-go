package arr

// Two pointer
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	var quickSort func(start, end int)
	quickSort = func(start, end int) {
		if start >= end {
			return
		}

		pivot := nums[start]
		lp, rp := start, end

		for lp < rp {
			for lp < rp && nums[rp] >= pivot {
				rp--
			}
			for lp < rp && nums[lp] <= pivot {
				lp++
			}
			nums[lp], nums[rp] = nums[rp], nums[lp]
		}
		nums[start], nums[lp] = nums[lp], nums[start]

		quickSort(start, lp-1)
		quickSort(lp+1, end)
	}

	quickSort(0, n-1)

	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		lp, rp := i+1, n-1

		for lp < rp {
			if nums[lp]+nums[rp] < target {
				for lp < rp && nums[lp+1] == nums[lp] {
					lp++
				}
				lp++
				continue
			} 
			
			if nums[lp]+nums[rp] > target {
				for lp < rp && nums[rp-1] == nums[rp] {
					rp--
				}
				rp--
				continue
			}

			res = append(res, []int{nums[i], nums[lp], nums[rp]})

			for lp < rp && nums[lp] == nums[lp+1] {
				lp++
			}
			for lp < rp && nums[rp] == nums[rp-1] {
				rp--
			}

			lp++
			rp--
		}
	}

	return res
}
