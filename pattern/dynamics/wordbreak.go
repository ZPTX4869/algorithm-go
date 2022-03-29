package dynamics

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)

	wordMap := make(map[string]bool, 0)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 检查s[j:i]是否存在于wordMap中
			_, ok := wordMap[s[j:i]]

			if (j == 0 || dp[j]) && ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
