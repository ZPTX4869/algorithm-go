package sort

func isStraight(nums []int) bool {
	var quickSort func(start, end int)
	quickSort = func(start, end int) {
		if start >= end {
			return
		}

		pivot := nums[start]
		left, right := start, end
		for left < right {
			for left < right && nums[right] >= pivot {
				right--
			}
			for left < right && nums[left] <= pivot {
				left++
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
		nums[start], nums[left] = nums[left], nums[start]

		quickSort(start, left-1)
		quickSort(left+1, end)
	}

	quickSort(0, len(nums)-1)

	zeroCnt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroCnt += 1
			continue
		}

		if nums[i] == nums[i+1] {
			return false
		}

		if nums[i+1]-nums[i] != 1 {
			balance := nums[i+1] - nums[i] - 1
			zeroCnt -= balance
			if zeroCnt < 0 {
				return false
			}
		}
	}

	return true
}

func isStraight2(nums []int) bool {
	used := make(map[int]bool, 0)
	max, min := 0, 14

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		if _, ok := used[nums[i]]; ok {
			return false
		}
		used[nums[i]] = true

		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	return max-min < 5
}
