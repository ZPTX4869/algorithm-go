package dp

func isMatch(s string, p string) bool {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(p))
		for j:= range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if i == -1 {
			if j == -1 {
				return true
			}
			if p[j] == '*' {
				return dp(i, j-2)
			}

			return false
		}

		if j == -1 {
			return false
		}

		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		if s[i] == p[j] || p[j] == '.' {
			if dp(i-1, j-1) {
				memo[i][j] = 1
			} else {
				memo[i][j] = 0
			}
		}

		if p[j] == '*' {
			if j > 0 && (p[j-1] == s[i] || p[j-1] == '.') {
				if dp(i-1, j) || dp(i-1, j-2) || dp(i, j-2) {
					memo[i][j] = 1
				} else {
					memo[i][j] = 0
				}
			} else {
				if dp(i, j-2) {
					memo[i][j] = 1
				} else {
					memo[i][j] = 0
				}
			}
		}

		return memo[i][j] == 1
	}

	return dp(len(s)-1, len(p)-1)
}
