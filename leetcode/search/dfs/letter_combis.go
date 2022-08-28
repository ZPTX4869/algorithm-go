package dfs

var buttonMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	var combis string

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(combis) == len(digits) {
			if combis != "" {
				res = append(res, combis)
			}
			return
		}

		letters, ok := buttonMap[digits[idx]]
		if !ok {
			return
		}

		for i := 0; i < len(letters); i++ {
			combis += letters[i]
			backtrack(idx + 1)
			combis = combis[:len(combis)-1]
		}
	}

	backtrack(0)

	return res
}
