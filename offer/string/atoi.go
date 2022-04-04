package string

import (
	"math"
)

func myAtoi(s string) int {
	nums := make([]int64, 0)
	neg := false

	i, j := 0, 0
	for ; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if s[i] == '-' {
				neg = true
			}
			j = i + 1
			break
		}

		if s[i] >= '0' && s[i] <= '9' {
			j = i
			break
		}

		if s[i] != ' ' {
			return 0
		}
	}

	for ; j < len(s); j++ {
		if s[j] < '0' || s[j] > '9' {
			break
		}
		nums = append(nums, int64(s[j]-'0'))
	}

	return calcNums(nums, neg)
}

func calcNums(nums []int64, neg bool) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	var res int64

	for i := 0; i < n; i++ {
		res = res*10 + nums[i]

		if res > int64(math.MaxInt32) {
			if neg == true {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	if neg == true {
		res = -res
	}

	return int(int32(res))
}
