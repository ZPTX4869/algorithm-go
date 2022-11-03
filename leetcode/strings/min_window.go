package str

import (
	"math"
)

func minWindow(s string, t string) string {
	var res string
	need := make(map[byte]int, len(s))
	window := make(map[byte]int, len(t))
	length := math.MaxInt32

	for _, v := range t {
		need[byte(v)] += 1
	}

	valid := 0
	lp, rp := 0, 0
	for rp < len(s) {
		c := s[rp]
		rp++

		if need[c] > 0 {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}

		for lp <= rp && valid == len(need) {
			if rp-lp < length {
				res = s[lp:rp]
				length = rp - lp
			}

			c = s[lp]
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

	return res
}
