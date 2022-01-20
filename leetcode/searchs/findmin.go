package searchs

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
	}

	return nums[left]
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == nums[right] {
			if searchRight(nums, mid, right) {
				left = mid + 1
			} else {
				right = mid
			}
		} else if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
	}

	return nums[left]
}

func searchRight(nums []int, start, end int) bool {
	for cur := start; cur < end; cur++ {
		if nums[cur] < nums[start] {
			return true
		}
	}
	return false
}
