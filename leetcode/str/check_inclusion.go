package str

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, len(s1))
	window := make(map[byte]int, len(s1))

	for _, v := range s1 {
		need[byte(v)] += 1
	}

	valid := 0
	lp, rp := 0, 0
	for rp < len(s2) {
		c := s2[rp]
		rp++

		if need[c] > 0 {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}

		for lp < rp && valid == len(need) {
			if rp-lp == len(s1) {
				return true
			}

			c = s2[lp]
			lp++

			if need[c] > 0 {
				window[c] -= 1
				if window[c] < need[c] {
					valid -= 1
					break
				}
			}
		}
	}

	return false
}
