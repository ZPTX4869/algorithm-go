package array

func exchange(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	res := make([]int, n)
	left, right := 0, n-1

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			res[right] = nums[i]
			right -= 1
		} else {
			res[left] = nums[i]
			left += 1
		}
	}

	return res
}

func exchange2(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	left, right := 0, n-1
	for left != right {
		if nums[left]%2 != 0 {
			left += 1
			continue
		}
		if nums[right]%2 == 0 {
			right -= 1
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}