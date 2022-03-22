package sorts

import (
	"strconv"
)

func minNumber(nums []int) string {
	res := ""

	numStrs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStrs[i] = strconv.Itoa(nums[i])
	}

	larger := func(numStr1, numStr2 string) bool {
		num1, _ := strconv.Atoi(numStr1 + numStr2)
		num2, _ := strconv.Atoi(numStr2 + numStr1)
		return num1 > num2
	}

	var quickSort func(start, end int)
	quickSort = func(start, end int) {
		if start >= end {
			return
		}

		pivot := numStrs[start]
		l, r := start, end
		for l < r {
			for l < r && larger(numStrs[r], pivot) {
				r--
			}
			for l < r && !larger(numStrs[l], pivot) {
				l++
			}
			numStrs[l], numStrs[r] = numStrs[r], numStrs[l]
		}
		numStrs[start], numStrs[l] = numStrs[l], numStrs[start]

		quickSort(start, l-1)
		quickSort(l+1, end)
	}

	quickSort(0, len(numStrs)-1)

	for _, numStr := range numStrs {
		res += numStr
	}

	return string(res)
}
