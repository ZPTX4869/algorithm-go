package dp

import "algorithm-go/util"

func maxProfit(prices []int) int {
	// dp0和dp1分别表示空仓和满仓状态
	dp0 := make([]int, len(prices))
	dp1 := make([]int, len(prices))

	dp0[0] = 0
	dp1[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp0[i] = util.Max(dp0[i-1], dp1[i-1]+prices[i])
		dp1[i] = util.Max(dp1[i-1], -prices[i])
	}

	return dp0[len(prices)-1]
}
