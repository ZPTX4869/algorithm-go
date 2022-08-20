package stack

func calculate(s string) int {
	var helper func(bytes *[]byte) int
	helper = func(bytes *[]byte) int {
		stk := make([]int, 0)
		num := 0
		opt := byte('+')

		for len(*bytes) > 0 {
			c := (*bytes)[0]
			*bytes = (*bytes)[1:]

			if c == '(' {
				num = helper(bytes)
			}

			if isDigit(c) {
				num = 10*num + int(c-'0')
			}

			if (!isDigit(c) && c != ' ') || len(*bytes) == 0 {
				switch opt {
				case '+':
					stk = append(stk, num)
				case '-':
					stk = append(stk, -num)
				case '*':
					pre := stk[len(stk)-1]
					stk = stk[:len(stk)-1]
					stk = append(stk, pre*num)
				case '/':
					pre := stk[len(stk)-1]
					stk = stk[:len(stk)-1]
					stk = append(stk, pre/num)
				}

				opt = c
				num = 0
			}

			if c == ')' {
				break
			}
		}

		res := 0
		for len(stk) > 0 {
			res += stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}

		return res
	}

	bytes := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		bytes = append(bytes, s[i])
	}

	return helper(&bytes)
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
