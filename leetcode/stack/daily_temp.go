package stack

func dailyTemperatures(temperatures []int) []int {
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
