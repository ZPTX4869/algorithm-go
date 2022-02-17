package dynamics

// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	maxLen := 1
	begin := 0

	// 初始化二维数组
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 每个单字母都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for l := 2; l < n+1; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}

			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				begin = i
				maxLen = l
			}
		}
	}

	return s[begin : begin+maxLen]
}

// 暴力求解
// func longestPalindrome(s string) string {
// 	size := len(s)
// 	longest := ""

// 	for i := 0; i < size; i++ {
// 		for j := size - 1; j >= i; j-- {
// 			if isPalindrome(s, i, j) {
// 				if len(longest) < j-i+1 {
// 					longest = s[i : j+1]
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return longest
// }

func partition(s string) [][]string {
	var res [][]string

	for i := 0; i < len(s); i++ {
	}

	return res
}

func minCut(s string) int {
	return 0
}

// func isPalindrome(s string, left, right int) bool {
// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}

// 	return true
// }
