package stacks

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			size := len(stack)
			num1, num2 := stack[size-1], stack[size-2]
			stack = stack[:size-2]

			switch token {
			case "+":
				stack = append(stack, num2+num1)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			default:
				stack = append(stack, num2/num1)
			}
		}
	}

	return stack[0]
}
