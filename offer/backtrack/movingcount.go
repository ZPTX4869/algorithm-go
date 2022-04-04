package backtrack

type pair struct {
	x int
	y int
}

var directions = []pair{
	{+1, 0},
	{0, +1},
}

func movingCount(m int, n int, k int) int {
	var res int
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var backtrack func(row, col int)
	backtrack = func(row, col int) {
		if visited[row][col] || sumDigits(row)+sumDigits(col) > k {
			return
		}

		res += 1
		visited[row][col] = true

		for _, p := range directions {
			nextRow, nextCol := row+p.x, col+p.y
			if nextRow < m && nextCol < n {
				backtrack(nextRow, nextCol)
			}
		}
	}

	backtrack(0, 0)

	return res
}

func movingCount2(m int, n int, k int) int {
	queue := make([]pair, 0)
	visited := make(map[pair]bool)

	queue = append(queue, pair{0, 0})
	
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if visited[p] || sumDigits(p.x) + sumDigits(p.y) > k {
			continue
		}
		visited[p] = true
		if p.x + 1 < m {
			queue = append(queue, pair{p.x+1, p.y})
		}
		if p.y + 1 < n {
			queue = append(queue, pair{p.x, p.y+1})
		}
	}

	return len(visited)
}

func sumDigits(num int) int {
	var res int

	for num != 0 {
		digit := num % 10
		num /= 10
		res += digit
	}

	return res
}
