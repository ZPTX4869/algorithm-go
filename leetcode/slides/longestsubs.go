package slides

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 1
	l, r := 0, 1

	m := make(map[byte]int, 0)

	for r <= len(s) {
		if loc, ok := m[s[r-1]]; ok {
			l = loc + 1
			for k, v := range m {
				if v <= loc {
					delete(m, k)
				}
			}
		} 
		m[s[r-1]] = r - 1

		if r-l > maxLen {
			maxLen = r - l
		}

		fmt.Println(m)
		fmt.Printf("%d:%d %s\n", l, r, s[l:r])

		r += 1
	}

	return maxLen
}
