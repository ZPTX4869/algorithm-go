package str

// Leetcode: https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	max := 0
	window := make([]byte, 0)
	record := make(map[byte]bool, 0)

	lp, rp := 0, 0
	for rp < len(s) {
		c := s[rp]
		window = append(window, c)
		rp++

		if _, ok := record[c]; ok {
			for lp <= rp && s[lp] != c {
				delete(record, s[lp])
				window = window[1:]
				lp++
			}
			window = window[1:]
			lp++
		} else {
			record[c] = true
		}

		if len(window) > max {
			max = len(window)
		}
	}

	return max
}
