package dynamic

func maxProfit(prices []int) int {
	getMax := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	n := len(prices)
	if n < 2 {
		return 0
	}

	min := prices[0]
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = getMax(dp[i-1], prices[i]-min)

		if prices[i] < min {
			min = prices[i]
		}
	}

	return dp[n-1]
}

func maxProfit2(prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	max := 0
	stack := make([]int, 0)
	stack = append(stack, prices[0])
	prices = append(prices, 0)

	for i := 1; i < n+1; i++ {
		if prices[i] < stack[len(stack)-1] {
			for len(stack) > 1 && prices[i] < stack[len(stack)-1] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				profit := top - stack[0]
				if profit > max {
					max = profit
				}
			}
		}

		if prices[i] < stack[0] {
			stack = stack[1:]
		}

		stack = append(stack, prices[i])
	}

	return max
}
