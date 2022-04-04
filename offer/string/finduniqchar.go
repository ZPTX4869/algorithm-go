package string

func firstUniqChar(s string) byte {
	used := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		used[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if used[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

func firstUniqChar2(s string) byte {
	used := make([]int, 26)

	for i := 0; i < len(s); i++ {
		used[s[i]-'a'] += 1
	}

	for i := 0; i < len(s); i++ {
		if used[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}