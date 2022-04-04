package dynamic

import (
	"math"
)

// 硬币找零：https://leetcode-cn.com/problems/coin-change/
// 使用限定面额的硬币凑出目标数额，返回最少需要的硬币个数，如果无法凑整则返回-1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	n := len(coins)
	dp := make([]int, amount+1)

	coinMap := make(map[int]bool, n)
	for _, c := range coins {
		coinMap[c] = true
	}

	for i := 1; i < amount+1; i++ {
		min := math.MaxInt
		for _, c:= range coins {
			if i >= c && dp[i-c] != -1 {
				cnt := dp[i-c] + 1
				if cnt < min {
					min = cnt
				}
			}
		}

		if min == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = min
		}
	}

	return dp[amount]
}
