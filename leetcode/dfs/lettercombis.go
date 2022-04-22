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

// func letterCombinations(digits string) []string {
// 	res := make([]string, 0)

// 	for i := 0; i < len(digits); i++ {
// 		letters := make([]string, 0)

// 		if strs, ok := buttonMap[digits[i]]; ok {
// 			letters = append(letters, strs...)
// 		}

// 		if len(res) == 0 {
// 			res = append(res, letters...)
// 			continue
// 		}

// 		n := len(res)
// 		for i := 0; i < n; i++ {
// 			for j := 0; j < len(letters); j++ {
// 				res = append(res, res[i]+letters[j])
// 			}
// 		}
// 		res = res[n:]
// 	}

// 	return res
// }
