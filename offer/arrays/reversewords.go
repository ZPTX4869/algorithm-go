package arrays

func reverseWords(s string) string {
	removeEndSpace := func(s string) string {
		for s[0] == ' ' {
			s = s[1:]
		}
		for s[len(s)-1] == ' ' {
			s = s[0 : len(s)-1]
		}

		return s
	}

	s = removeEndSpace(s)

	n := len(s)
	if n == 0 {
		return ""
	}

	left, right := 0, n-1
	for right != left {
		if s[right] == ' ' && right == n-1 {
			s = s[:n-1]
			n = n - 1
			right -= 1
		} else if s[right] == ' ' {
			s = s[:left] + s[right:] + s[left:right]
			left = left + n - right
			right = n - 1
		} else {
			right -= 1
		}
	}

	s = s[0:left] + " " + s[left:]
	s = removeEndSpace(s)

	return s
}

func reverseWords2(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	if s[0] != ' ' {
		s = " " + s
		n += 1
	}

	res := make([]byte, 0)
	left, right := n-1, n
	for left >= 0 {
		if s[left] == ' ' {
			if right-left == 1 {
				left, right = left-1, right-1
			} else {
				res = append(res, s[left+1:right]...)
				res = append(res, ' ')
				left, right = left-1, left
			}
		} else {
			left -= 1
		}
	}

	for res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return string(res)
}
