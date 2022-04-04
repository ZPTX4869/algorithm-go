package backtrack

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	idx := 0

	used := make([][]bool, m)
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, n)
	}

	var backtrack func(row, col int) bool
	backtrack = func(row, col int) bool {
		if used[row][col] || board[row][col] != word[idx] {
			return false
		}

		if idx == len(word)-1 {
			return true
		}

		idx, used[row][col] = idx-1, true
		defer func() {
			idx, used[row][col] = idx-1, false
		}()

		if row-1 >= 0 {
			if backtrack(row-1, col) {
				return true
			}
		}
		if row+1 < m {
			if backtrack(row+1, col) {
				return true
			}
		}
		if col-1 >= 0 {
			if backtrack(row, col-1) {
				return true
			}
		}
		if col+1 < n {
			if backtrack(row, col+1) {
				return true
			}
		}

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j) {
				return true
			}
		}
	}

	return false
}
