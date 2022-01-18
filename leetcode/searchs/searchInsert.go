package searchs

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	return left
}
