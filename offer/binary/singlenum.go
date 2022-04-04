package binary

import "fmt"

// 除了一个数字外其他重复3次
func singleNumber(nums []int) int {
	var quickSort func(start, end int)
	quickSort = func(start, end int) {
		if start >= end {
			return
		}

		pivot := nums[start]
		l, r := start, end
		for l < r {
			for l < r && nums[r] >= pivot {
				r--
			}
			for l < r && nums[l] <= pivot {
				l++
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
		nums[start], nums[l] = nums[l], nums[start]

		quickSort(start, l)
		quickSort(l+1, end)
	}

	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	quickSort(0, n-1)
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			return nums[i]
		}
	}

	return nums[n-1]
}