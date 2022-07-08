package str

// 最长回文子串：https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	res := string(s[0])

	for i := 1; i < len(s); i++ {
		single := spread(s, i, i)
		double := spread(s, i-1, i)

		if len(single) > len(res) {
			res = single
		}
		if len(double) > len(res) {
			res = double
		}
	}

	return res
}

func spread(s string, lp, rp int) string {
	i, j := lp, rp
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i -= 1
		j += 1
	}

	return s[i+1 : j]
}
