package stacks

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		var result int

		if token == "+" {
			num1, num2 := pop2(stack)
			result = num2 + num1
		} else if token == "-" {
			num1, num2 := pop2(stack)
			result = num2 - num1
		} else if token == "*" {
			num1, num2 := pop2(stack)
			result = num2 * num1
		} else if token == "/" {
			num1, num2 := pop2(stack)
			result = num2 / num1
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				log.Fatalf("Error: %v occur when convert string to integer", err)
			}
			stack = append(stack, num)
			continue
		}

		stack = stack[:len(stack)-2]
		stack = append(stack, result)
	}

	return stack[0]
}

func pop2(nums []int) (int, int) {
	size := len(nums)

	num1 := nums[size-1]
	num2 := nums[size-2]

	return num1, num2
}
