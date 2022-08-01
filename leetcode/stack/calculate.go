package stack

func calculate(s string) int {
	return eval(s, 0, len(s)-1)
}

func calculate_(s string) int {
	return eval(s, 0, len(s)-1)
}

func eval(s string, left, right int) int {
	num := 0
	opt := byte('+')
	stk := make([]int, 0)

	for i := left; i <= right; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = 10*num + int(s[i]-'0')
		}

		if s[i] == '(' {
			layer := 0
			j := i
			for j <= right {
				if s[j] == '(' {
					layer += 1
				}
				if s[j] == ')' {
					layer -= 1
				}
				if layer == 0 {
					break
				}
				j++
			}
			num = eval(s, i+1, j-1) // 递归计算字表达式
			i = j
		}

		if s[i] < '0' || s[i] > '9' || i == right {
			switch opt {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			case '*':
				stk[len(stk)-1] *= num
			case '/':
				stk[len(stk)-1] /= num
			}
			opt = s[i]
			num = 0
		}
	}

	res := 0
	for _, v := range stk {
		res += v
	}

	return res
}
