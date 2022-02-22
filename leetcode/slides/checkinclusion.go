package slides

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)

	cnt1, cnt2 := make([]int, 26), make([]int, 26)
	for i := 0; i < m; i++ {
		cnt1[s1[i]-'a'] += 1
		cnt2[s2[i]-'a'] += 1
	}

	for i := m; i < n; i++ {
		if isEqual(cnt1, cnt2) {
			return true
		}

		cnt2[s2[i-m]-'a'] -= 1
		cnt2[s2[i]-'a'] += 1
	}

	return isEqual(cnt1, cnt2)
}

func isEqual(cnt1, cnt2 []int) bool {
	for i := 0; i < len(cnt1); i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}

	return true
}
