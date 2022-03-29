package queues

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	m := len(mat)
	n := len(mat[0])

	distMat := make([][]int, 0)
	visited := make([][]int, 0)
	for i := 0; i < m; i++ {
		distMat = append(distMat, make([]int, n))
		visited = append(visited, make([]int, n))
	}

	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
				visited[i][j] = 1
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]

			for j := 0; j < 4; j++ {
				newX, newY := x+dx[j], y+dy[j]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && visited[newX][newY] == 0 {
					distMat[newX][newY] = distMat[x][y] + 1
					visited[newX][newY] = 1
					queue = append(queue, []int{newX, newY})
				}
			}
		}
	}

	return distMat
}
