package bin

func searchRange(nums []int, target int) []int {
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	lp, rp := 0, len(nums)
	for lp < rp {
		mid := lp + (rp-lp)/2

		if target == nums[mid] {
			rp = mid
		} else if target < nums[mid] {
			rp = mid
		} else if target > nums[mid] {
			lp = mid + 1
		}
	}

	if nums[lp] != target {
		return -1
	}

	return lp
}

func searchRight(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	lp, rp := 0, len(nums)
	for lp < rp {
		mid := lp + (rp-lp)/2

		if target == nums[mid] {
			lp = mid + 1
		} else if target < nums[mid] {
			rp = mid
		} else if target > nums[mid] {
			lp = mid + 1
		}
	}

	if nums[rp-1] != target {
		return -1
	}

	return rp - 1
}
