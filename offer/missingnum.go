package offer

func missingNumber(nums []int) int {
	var res int
	var binarySearch func(s, e int)

	binarySearch = func(s, e int) {
		left := (e-s)/2 + s
		right := left + 1

		if nums[right] - nums[left] > 1 {
			res = left+1
			return
		}

		if (nums[left] - nums[s]) != (left - s) {
			binarySearch(s, left)
		} else if (nums[e]- nums[right]) != (e - right) {
			binarySearch(right, e)
		}
	}

	binarySearch(0, len(nums))

	return res
}