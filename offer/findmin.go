package offer

func findMin(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := (end-start)/2 + start

		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end -= 1
		}
	}

	return nums[start]
}
