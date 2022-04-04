package stack

func validateStackSequences(pushed []int, popped []int) bool {
	var res bool
	stack := make([]int, 0)

	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && popped[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}

	if len(stack) == 0 {
		res = true
	}

	return res
}
