package search

func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if mid == 0 {
			if nums[0] == target {
				return 0
			}
			return -1
		}

		if nums[mid-1] >= nums[left] {
			if searchSorted(nums, left, mid, target) {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if searchSorted(nums, mid, right, target) {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return -1
}

func search2(nums []int, target int) bool {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		if mid == 0 {
			if nums[0] == target {
				return true
			}
			return false
		}

		if nums[mid-1] > nums[left] || (nums[mid-1] == nums[left] && isAllEqual(nums, left, mid-1, nums[left])) {
			if searchSorted(nums, left, mid, target) {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if searchSorted(nums, mid, right, target) {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return false
}

func searchSorted(nums []int, start, end, target int) bool {
	if target >= nums[start] && target <= nums[end-1] {
		return true
	}

	return false
}

func isAllEqual(nums []int, start, end, repeat int) bool {
	res := true
	for i := start; i < end; i++ {
		if nums[i] != repeat {
			res = false
			break
		}
	}

	return res
}
