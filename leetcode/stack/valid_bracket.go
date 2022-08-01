package stack

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}

	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stk := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := match[s[i]]; ok {
			if len(stk) == 0 || match[s[i]] != stk[len(stk)-1] {
				return false
			}
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, s[i])
		}
	}

	return len(stk) == 0
}