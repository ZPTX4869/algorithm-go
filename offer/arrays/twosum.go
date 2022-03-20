package arrays

func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	left, right := 0, n-1
	for left != right {
		sum := nums[left] + nums[right]
		if target < sum {
			right -= 1
		} else if target > sum {
			left += 1
		} else {
			return []int{nums[left], nums[right]}
		}
	}

	return nil
}
