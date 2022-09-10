package dp

import (
	"algorithm-go/util/maths"
	"math"
)

// K站中转最便宜的航班：https://leetcode.cn/problems/cheapest-flights-within-k-stops/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+2)
	}

	graph := make(map[int][][]int, 0)
	for _, v := range flights {
		from, to, weight := v[0], v[1], v[2]

		if _, ok := graph[to]; !ok {
			graph[to] = make([][]int, 0)
		}

		graph[to] = append(graph[to], []int{from, weight})
	}

	var minPath func(dst, step int) int
	minPath = func(dst, step int) int {
		if dst == src {
			return 0
		}

		if step == 0 {
			return -1
		}

		if memo[dst][step] != 0 {
			return memo[dst][step]
		}

		res := math.MaxInt32
		for _, v := range graph[dst] {
			from, weight := v[0], v[1]

			subRes := minPath(from, step-1)
			if subRes != -1 {
				res = maths.Min(res, subRes+weight)
			}
		}

		if res == math.MaxInt32 {
			memo[dst][step] = -1
		} else {
			memo[dst][step] = res
		}

		return memo[dst][step]
	}

	return minPath(dst, k+1)
}

func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	// At most k stops means we can move at most k+1 steps.
	step := k + 1

	dp := make([][]int, step+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0

	for i := 1; i <= step; i++ {
		for _, v := range flights {
			from, to, weight := v[0], v[1], v[2]
			dp[i][to] = maths.Min(dp[i][to], dp[i-1][from]+weight)
		}
	}

	res := inf
	for i := 0; i <= step; i++ {
		res = maths.Min(res, dp[i][dst])
	}

	if res == inf {
		return -1
	}

	return res
}
