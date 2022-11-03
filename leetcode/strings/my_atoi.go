package str

import "math"

func myAtoi(s string) int {
	var res int64
	var neg bool
	var i int

	for i < len(s) && s[i] == ' ' {
		i += 1
	}

	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			neg = true
		}
		i += 1
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = 10*res + int64(s[i]-'0')
		if !neg && res > int64(math.MaxInt32) {
			return math.MaxInt32
		}
		if neg && -res <= int64(math.MinInt32) {
			return math.MinInt32
		}
		i += 1
	}

	if neg {
		res = -res
	}

	return int(res)
}
