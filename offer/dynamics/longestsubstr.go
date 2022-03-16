package dynamics

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	used := make(map[byte]bool, 0)
	dp := make([]int, n+1)
	dp[0] = 0

	max, start := 0, 0
	for i := 0; i < n; i++ {
		if _, ok := used[s[i]]; !ok {
			dp[i+1] = dp[i] + 1
			used[s[i]] = true
		} else {
			for j := start; j < i; j++ {
				if s[j] == s[i] {
					start = j + 1
					break
				}
				used[s[j]] = false
			}
			dp[i+1] = i - start + 1
		}

		if dp[i+1] > max {
			max = dp[i+1]
		}
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	var res int

	n := len(s)
	if n == 0 {
		return res
	}

	used := make(map[byte]bool, 0)
	left, right := 0, 0

	for right < n {
		if _, ok := used[s[right]]; !ok {
			used[s[right]] = true
		} else {
			for s[left] != s[right] {
				delete(used, s[left])
				left += 1
			}
			left = left + 1
		}

		if right-left+1 > res {
			res = right - left + 1
		}

		right += 1
	}

	return res
}
