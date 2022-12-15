package stack

// 每日温度：https://leetcode.cn/problems/daily-temperatures/

// 暴力枚举
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		nxt := i + 1
		for nxt < n && temperatures[nxt] <= temperatures[i] {
			nxt++
		}

		if nxt == n {
			ans[i] = 0
		} else {
			ans[i] = nxt - i
		}
	}

	return ans
}

// 正向单调栈
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stk := []int{0}

	for i := 1; i < n; i++ {
		if temperatures[i] > temperatures[stk[len(stk)-1]] {
			for len(stk) > 0 && temperatures[i] > temperatures[stk[len(stk)-1]] {
				top := stk[len(stk)-1]
				ans[top] = i - top
				stk = stk[:len(stk)-1]
			}
		}

		stk = append(stk, i)
	}

	for i := 0; i < len(stk); i++ {
		ans[stk[i]] = 0
	}

	return ans
}

// 逆向单调栈
func dailyTemperatures3(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stk := make([]int, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] <= temperatures[i] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) > 0 {
			res[i] = stk[len(stk)-1] - i
		}

		stk = append(stk, i)
	}

	return res
}
