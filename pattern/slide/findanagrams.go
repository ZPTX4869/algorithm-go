package slide

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	
	m, n := len(s), len(p)
	if m < n {
		return result
	}

	cnt1, cnt2 := make([]int, 26), make([]int, 26)
	for i := 0; i < n; i++ {
		cnt1[p[i]-'a'] += 1
		cnt2[s[i]-'a'] += 1
	}

	for i := n; i < m; i++ {
		if isEqual(cnt1, cnt2) {
			result = append(result, i-n)
		}

		cnt2[s[i-n]-'a'] -= 1
		cnt2[s[i]-'a'] += 1
	}

	if isEqual(cnt1, cnt2) {
		result = append(result, m-n)
	}

	return result
}