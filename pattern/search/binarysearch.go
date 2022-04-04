package search

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// 不要求和后除2，可能发生越界
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func binarySearch2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		// 不要求和后除2，可能发生越界
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left == right && nums[left] == target {
		return left
	}

	return -1
}

func binarySearch3(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left+1 < right {
		// 不要求和后除2，可能发生越界
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	}

	return -1
}
